# sns

UI branch: https://github.com/ideadevice/calm-ui/tree/hack/sns


## Description

SNS is many to many pub/sub messaging service that is highly available, durable, secure.
It can be used for event notification, monitoring applications, workflow systems and time-sensitive information updates.
Provides significant advantages to developers who build applications that rely on real-time events


## UsesCases
1. Effective communication between nutanix services.
  - Nutanix providers On prem features like budget utilizations. Currently when customers create VMs through Prism or deploy applications through CALM budget updates are done at end of the day for all the resource utilization. This is one big job mainted by team BEAM running at end of everyday. Suppose there is an event named VM_Create which is triggered by prism and calm as soon as resource allocation is done and BEAM being subscriber of it. Prism and CALM publishes event with metdata about the VM and beam updates budget utilization for the category to which this VM belongs in real team. This is avoids huge cron based jobs and updates in real-time.
  - There is a service named Curator maintained by team Stargate to free disk space when reference count is zero. This job also runs periodically. Suppose there is a event Disk_space_reference_zero and it is triggered by services while deleting vDisk and reference count is zero, disk space can be made available in the real time.
 
2. SNS..aaS
  - Nutanix is providing mutiple features/services for On prem customers. Why not provide this SNS? We do have customers who are crowd facing and they would like to send mobile/email notification to their customers as part of festival campaign. Why go to some other cloud service just for this?
  
 
