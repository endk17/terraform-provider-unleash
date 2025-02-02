---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "unleash_feature_v2 Resource - terraform-provider-unleash"
subcategory: ""
description: |-
  (Experimental) Provides a resource for managing unleash features with variants and environment strategies all in a single resource.
---

# unleash_feature_v2 (Resource)

(Experimental) Provides a resource for managing unleash features with variants and environment strategies all in a single resource.

## Example Usage

```terraform
resource "unleash_feature_v2" "with_env_strategies" {
  name               = "my_nice_feature"
  description        = "manages my nice feature"
  type               = "release"
  project_id         = "default"
  archive_on_destroy = false

  environment {
    name    = "production"
    enabled = false
  }

  environment {
    name    = "development"
    enabled = true

    strategy {
      name = "remoteAddress"
      parameters = {
        IPs = "189.434.777.123,host.test.com"
      }
    }
    strategy {
      name = "flexibleRollout"
      parameters = {
        rollout    = "68"
        stickiness = "random"
        groupId    = "toggle"
      }
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Feature name
- `project_id` (String) The feature will be created in the given project
- `type` (String) Feature type

### Optional

- `archive_on_destroy` (Boolean) Whether to archive the feature toggle on destroy. Default is `true`. When `false`, it will permanently delete the feature toggle.
- `description` (String) Feature description
- `environment` (Block List) Use this to enable a feature in an environment and add strategies (see [below for nested schema](#nestedblock--environment))
- `variant` (Block List) Feature variant (see [below for nested schema](#nestedblock--variant))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--environment"></a>
### Nested Schema for `environment`

Required:

- `name` (String) Environment name

Optional:

- `enabled` (Boolean) Whether the feature is on/off in the environment. Default is `true` (on)
- `strategy` (Block Set) Strategy to add in the environment (see [below for nested schema](#nestedblock--environment--strategy))

<a id="nestedblock--environment--strategy"></a>
### Nested Schema for `environment.strategy`

Required:

- `name` (String) Strategy unique name

Optional:

- `parameters` (Map of String) Strategy parameters. All the values need to informed as strings.

Read-Only:

- `id` (String) Strategy ID



<a id="nestedblock--variant"></a>
### Nested Schema for `variant`

Required:

- `name` (String) Variant name

Optional:

- `overrides` (Block Set) Overrides existing context field values. Values are comma separated e.g `v1, v2, ...`) (see [below for nested schema](#nestedblock--variant--overrides))
- `payload` (Block Set, Max: 1) Variant payload. The type of the payload can be `string`, `json` or `csv` (see [below for nested schema](#nestedblock--variant--payload))
- `stickiness` (String) Variant stickiness. Default is `default`.
- `weight` (Number) Variant weight. Only considered when the `weight_type` is `fix`. It is calculated automatically if the `weight_type` is `variable`.
- `weight_type` (String) Variant weight type. The weight type can be `fix` or `variable`. Default is `variable`.

<a id="nestedblock--variant--overrides"></a>
### Nested Schema for `variant.overrides`

Required:

- `context_name` (String)
- `values` (List of String)


<a id="nestedblock--variant--payload"></a>
### Nested Schema for `variant.payload`

Required:

- `type` (String)
- `value` (String)


