#!/bin/sh

SEGMENT=192.168.42
WIFINAME=PlanB
PASSWORD=12345678

sudo apt-get -y install openssh-server hostapd isc-dhcp-server iptables-persistent

cat <<EOF >> /etc/dhcpcd.conf 
interface wlan0
static ip_address=$SEGMENT.1/24
nohook wpa_supplicant
EOF

cat > /etc/network/interfaces.d/wlan0 <<EOF
iface wlan0 inet static
address $SEGMENT.1
netmask 255.255.255.0
EOF

# Proxy DNS - Not used here
#sudo apt-get install dnsmasq
#cat > /etc/dnsmasq.conf <<EOF
#interface=eth0
#listen-address=192.168.1.1
#dhcp-range=192.168.1.50,192.168.1.100,12h
#server=8.8.8.8
#bind-interfaces
#domain-needed
#bogus-priv
#EOF

echo "net.ipv4.ip_forward=1" > /etc/sysctl.conf
echo 1 > /proc/sys/net/ipv4/ip_forward

#Won't work
cat <<EOF >> /etc/dhcp/dhcpd.conf 
subnet $SEGMENT.0 netmask 255.255.255.0 {
         range $SEGMENT.10 $SEGMENT.20;
         option broadcast-address $SEGMENT.255;
         option routers $SEGMENT.1;
         option domain-name "local";
         option domain-name-servers 8.8.8.8, 8.8.4.4;
 }
EOF

cat > /etc/hostapd/hostapd.conf << EOF
ssid=$WIFINAME
wpa_passphrase=$PASSWORD

country_code=US

interface=wlan0
driver=nl80211

wpa=2
wpa_key_mgmt=WPA-PSK
rsn_pairwise=CCMP

macaddr_acl=0
auth_algs=1

logger_syslog=1
logger_syslog_level=4
logger_stdout=-1
logger_stdout_level=0

wmm_enabled=0

channel=11
EOF

cat <<EOF >> /etc/default/isc-dhcp-server
INTERFACESv4="wlan0"
INTERFACESv6=""
EOF

sudo systemctl unmask hostapd
sudo systemctl enable hostapd
sudo systemctl start hostapd

sudo update-rc.d hostapd enable
sudo update-rc.d isc-dhcp-server enable

sudo rfkill block bluetooth
sudo rfkill unblock wifi

