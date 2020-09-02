package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/xackery/log"
)

// WeapData represents a XML payload of weapons
type WeapData struct {
	XMLName   xml.Name `xml:"xml"`
	Text      string   `xml:",chardata"`
	ItemsProp struct {
		Text  string `xml:",chardata"`
		Items []struct {
			Text       string `xml:",chardata"`
			Name       string `xml:"name,attr"`
			Properties []struct {
				Text       string `xml:",chardata"`
				Name       string `xml:"name,attr"`
				Value      string `xml:"value,attr"`
				Class      string `xml:"class,attr"`
				Properties []struct {
					Text       string `xml:",chardata"`
					Name       string `xml:"name,attr"`
					Value      string `xml:"value,attr"`
					Param1     string `xml:"param1,attr"`
					Class      string `xml:"class,attr"`
					Properties struct {
						Text   string `xml:",chardata"`
						Name   string `xml:"name,attr"`
						Value  string `xml:"value,attr"`
						Param1 string `xml:"param1,attr"`
					} `xml:"property"`
				} `xml:"property"`
			} `xml:"property"`
			EffectGroup []struct {
				Text           string `xml:",chardata"`
				Name           string `xml:"name,attr"`
				PassiveEffects []struct {
					Text         string `xml:",chardata"`
					Name         string `xml:"name,attr"`
					Operation    string `xml:"operation,attr"`
					Value        string `xml:"value,attr"`
					Tags         string `xml:"tags,attr"`
					Tier         string `xml:"tier,attr"`
					MatchAllTags string `xml:"match_all_tags,attr"`
				} `xml:"passive_effect"`
				DisplayValue struct {
					Text  string `xml:",chardata"`
					Name  string `xml:"name,attr"`
					Value string `xml:"value,attr"`
				} `xml:"display_value"`
			} `xml:"effect_group"`
		} `xml:"item"`
	} `xml:"items"`
}

func main() {
	log := log.New()
	err := run()
	if err != nil {
		log.Error().Err(err).Msg("failed")
	}
	log.Info().Msg("success")
}

func run() error {
	//log := log.New()
	if len(os.Args) < 3 {
		fmt.Printf("usage: %s <items.xml> <Localization.txt>\n", os.Args[0])
		os.Exit(1)
	}
	path := os.Args[1]
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	xmlData, err := ioutil.ReadAll(f)
	weap := WeapData{}
	xmlData = []byte("<xml>" + string(xmlData) + "</xml>")
	err = xml.Unmarshal(xmlData, &weap)
	if err != nil {
		return err
	}

	path = os.Args[2]
	fl, err := os.Open(path)
	if err != nil {
		return err
	}

	translations := make(map[string]string)
	r := bufio.NewScanner(fl)
	for r.Scan() {
		records := strings.Split(r.Text(), ",")
		if len(records) < 6 {
			continue
		}
		translations[records[0]] = records[5]
	}

	for _, item := range weap.ItemsProp.Items {

		maxRange := ""
		if strings.Index(item.Name, "ammo") == 0 {
			continue
		}
		if item.Name == "meleeHandZombie01" {
			continue
		}
		//if strings.Index(item.Name, "thrownAmmo") == 0 {
		//	continue
		//}
		//if strings.Index(item.Name, "thrownGrenade") == 0 {
		//	continue
		//}

		if strings.Index(item.Name, "ammo") == 0 || strings.Index(item.Name, "thrownAmmo") == 0 {
			continue
		}

		if item.Name == "resourceNail" {
			continue
		}

		var damage float64
		var apm float64

		for _, p := range item.Properties {
			if p.Class != "Action0" {
				continue
			}
			for _, p2 := range p.Properties {
				if p2.Name == "Delay" {
					apm, err = strconv.ParseFloat(p2.Value, 64)
					if err != nil {
						return fmt.Errorf("parse %s Delay %s: %w", item.Name, p2.Value, err)
					}
					apm *= 60
				}
			}
		}
		for _, eg := range item.EffectGroup {
			for _, pe := range eg.PassiveEffects {
				if item.Name == "thrownGrenade" {
					apm = 1.2 * 60
				}
				if pe.Name == "EntityDamage" && pe.Operation == "base_set" {
					damage, err = strconv.ParseFloat(pe.Value, 64)
					if err != nil {
						return fmt.Errorf("parse %s EntityDamage %s: %w", item.Name, pe.Value, err)
					}
				}
				if pe.Name == "AttacksPerMinute" && pe.Operation == "base_set" {
					apm, err = strconv.ParseFloat(pe.Value, 64)
					if err != nil {
						return fmt.Errorf("parse %s AttacksPerMinute %s: %w", item.Name, pe.Value, err)
					}
				}
				if pe.Name == "RoundsPerMinute" && pe.Operation == "base_set" {
					apm, err = strconv.ParseFloat(pe.Value, 64)
					if err != nil {
						return fmt.Errorf("parse %s RoundsPerMinute %s: %w", item.Name, pe.Value, err)
					}
				}

				if pe.Name == "MaxRange" && pe.Operation == "base_set" {
					maxRange = pe.Value
				}
			}
		}
		if damage == 0 {
			continue
		}
		name, ok := translations[item.Name]
		if !ok {
			name = item.Name
			//			log.Warn().Msgf("no translation for %s found", item.Name)
		}
		if maxRange == "" {
			maxRange = "5"
		}
		if damage == 0 || apm == 0 {
			fmt.Printf("%s|0|0|0|%s\n", name, maxRange)
		} else {
			fmt.Printf("%s|%0.1f|%0.1f|%0.1f|%s\n", name, float32(damage/(apm/60)), damage, float32(apm/60), maxRange)
		}
	}

	return nil
}
