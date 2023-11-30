#!/usr/bin/env bash

cp ./docs/swagger.json /tmp

git branch -D gh-pages
git checkout --orphan gh-pages
git rm --cached -r .
cat .gitignore | xargs rm -rf .gitignore .github .env*
rm -rf ./*
git rm -rf

wget --no-clobber https://github.com/swagger-api/swagger-ui/archive/refs/tags/v5.10.3.zip -O swag.zip

unzip swag.zip
mkdir public
mv swagger-ui-5.10.3/dist/* public
rm -rf swagger-ui-5.10.3

cp /tmp/swagger.json public
sed -i 's/https:\/\/petstore.swagger.io\/v2\/swagger.json/.\/swagger.json/' public/*

mv public/* .
rm -rf swag.zip public

git add .
git commit -s -m "deploy swagger"
git checkout dev
git push -f origin gh-pages
