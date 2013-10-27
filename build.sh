# Simple script to do
#  idl2go
#  compile

# valid for Ubunut anyway
GOROOT=/usr/lib/go/

# turn idl into go code
barrister readability.idl  | $GOROOT/bin/idl2go -i -p readability

# compile go src
gd -o readability-svr .




