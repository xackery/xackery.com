echo "Generating hugo"
hugo
echo "Syncing public/ with xackery.com..."
aws s3 sync public/ s3://xackery.com --acl public-read --region us-west-2 --delete --profile xackery