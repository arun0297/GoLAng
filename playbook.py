
"""
"""

import phantom.rules as phantom
import json
from datetime import datetime, timedelta

def on_start(container):
    phantom.debug('on_start() called')
    
    domain_reputation_2(container=container)
    geolocate_ip_1(container=container) 

    return
def geolocate_ip_1(action=None, success=None, container=None, results=None, handle=None, filtered_artifacts=None, filtered_results=None):
    
    phantom.debug('geolocate_ip_1() called')

    
    container_data = phantom.collect2(container=container, datapath=['artifact:*.cef.destinationAddress', 'artifact:*.id'])

    parameters = []
    
    
    for container_item in container_data:
        if container_item[0]:
            parameters.append({
                'ip': container_item[0],
                'context': {'artifact_id': container_item[1]},
            })

    phantom.act("geolocate ip", parameters=parameters, assets=['maxmind'],name="geolocate_ip_1") 
        
    
    return 
 


def domain_reputation_2(action=None, success=None, container=None, results=None, handle=None, filtered_artifacts=None, filtered_results=None):
    phantom.debug('domain_reputation_2() called')

    container_data = phantom.collect2(container=container, datapath=['artifact:*.cef.destinationDnsDomain', 'artifact:*.id'])

    parameters = []
    
    for container_item in container_data:
        if container_item[0]:
            parameters.append({
                'domain': container_item[0],
                'context': {'artifact_id': container_item[1]},
            })

    phantom.act("domain reputation", parameters=parameters, assets=['virustotal'], name="domain_reputation_2")

    return

def on_finish(container, summary):
    phantom.debug('on_finish() called')

    return
