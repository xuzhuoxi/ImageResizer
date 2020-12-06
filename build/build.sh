CRTDIR=$(cd `dirname $0`; pwd)

goxc -os="linux darwin windows freebsd openbsd" -arch="386 amd64 arm" -n=IconGen -pv=1.1 -wd=${CRTDIR}/../src -d=${CRTDIR}/./release -include=*.go,README*,LICENSE*