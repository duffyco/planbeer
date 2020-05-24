#!/bin/sh

iptables -F
iptables-restore < ./firewall.rules.intercept
