#!/bin/bash
scp root@ns1.vpsaddict.com:/etc/nsd/soh.re /tmp/
gohome /tmp/soh.re
scp /tmp/soh.re root@ns1.vpsaddict.com:/etc/nsd/soh.re
ssh root@ns1.vpsaddict.com "/etc/nsd/restart.sh"
