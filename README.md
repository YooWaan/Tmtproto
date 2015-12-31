Tmt: prototype code
========================

試作プロジェクト

Architecture
-------------------

<pre>
 +-----+ +------+ +------------+
 | SSL | | Auth | | FileSystem |
 +-----+ +------+ +------------+
 +-----------------------------+
 | Server:                     |
 +-----------------------------+
</pre>


How to development
----------------------

<pre>
# move to git clone dir.
cd ＜path to dir＞
# set GOPATH
export GOPATH=`pwd`

# setup env
./make.sh clean
# build
./make.sh build
# install
./make.sh install
</pre>
