---
layout: "google"
page_title: "Google: google_compute_project"
sidebar_current: "docs-google-datasource-compute-project"
description: |-
  Get the email address of the project's Google Cloud Storage service account
---

# google\_compute\_project

Use this data source to get the project detals

## Example Usage

```hcl
data "google_compute_project" "project" {}

output "project_number" {
  value = "${data.google_compute_project.project.project_number}"
} 
```

## Argument Reference

There are no arguments available for this data source.


## Attributes Reference

The following attributes are exported:

* `default_service_account` - Default service account used by VMs running in this project

* `project_number` - The number uniquely identifying the projec

* `self_link`- The URL for the resource
