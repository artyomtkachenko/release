# release
A simple web service to build and publish rpm packages 

# Should implement the following idea

First call:
curl -X POST -d '{rpm metadata} http://localhost/v1/init/rhel
Returns: { "Id": "uniq hash" }

Next call:
curl -X POST -F "file=@/tmp/package.zip" http://localhost/v1/build/rhel/${uniq hash}

