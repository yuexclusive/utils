#! /bin/bash

a=$(git tag)

git tag -d $a
git push -d origin $a