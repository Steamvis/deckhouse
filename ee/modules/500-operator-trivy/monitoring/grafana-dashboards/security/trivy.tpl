{
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": {
          "type": "prometheus",
          "uid": "${ds_prometheus}"
        },
        "enable": true,
        "hide": true,
        "iconColor": "rgba(0, 211, 255, 1)",
        "name": "Annotations & Alerts",
        "target": {
          "limit": 100,
          "matchAny": false,
          "tags": [],
          "type": "dashboard"
        },
        "type": "dashboard"
      }
    ]
  },
  "description": "A dashboard to visualize trivy image vulnerabilities using metrics from trivy-operator.",
  "editable": false,
  "fiscalYearStartMonth": 0,
  "gnetId": 16742,
  "graphTooltip": 0,
  "iteration": 1702933974309,
  "links": [],
  "liveNow": false,
  "panels": [
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${ds_prometheus}"
      },
      "fieldConfig": {
        "defaults": {
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          },
          "unit": "none"
        },
        "overrides": [
          {
            "matcher": {
              "id": "byName",
              "options": "Critical"
            },
            "properties": [
              {
                "id": "color",
                "value": {
                  "fixedColor": "red",
                  "mode": "fixed"
                }
              }
            ]
          },
          {
            "matcher": {
              "id": "byName",
              "options": "High"
            },
            "properties": [
              {
                "id": "color",
                "value": {
                  "fixedColor": "orange",
                  "mode": "fixed"
                }
              }
            ]
          },
          {
            "matcher": {
              "id": "byName",
              "options": "Medium"
            },
            "properties": [
              {
                "id": "color",
                "value": {
                  "fixedColor": "dark-yellow",
                  "mode": "fixed"
                }
              }
            ]
          },
          {
            "matcher": {
              "id": "byName",
              "options": "Low"
            },
            "properties": [
              {
                "id": "color",
                "value": {
                  "fixedColor": "green",
                  "mode": "fixed"
                }
              }
            ]
          },
          {
            "matcher": {
              "id": "byName",
              "options": "Unknown"
            },
            "properties": [
              {
                "id": "color",
                "value": {
                  "mode": "fixed"
                }
              }
            ]
          }
        ]
      },
      "gridPos": {
        "h": 9,
        "w": 24,
        "x": 0,
        "y": 0
      },
      "id": 2,
      "options": {
        "colorMode": "background",
        "graphMode": "none",
        "justifyMode": "center",
        "orientation": "auto",
        "reduceOptions": {
          "calcs": [
            "lastNotNull"
          ],
          "fields": "",
          "values": false
        },
        "textMode": "auto"
      },
      "pluginVersion": "8.5.13",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${ds_prometheus}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "sum(trivy_vulnerability_id{severity=\"Critical\",namespace=~\"$namespace\",vuln_id=~\"$vuln_id\",resource_name=~\"$resource_name\",resource=~\"$resource\"{{- range .Values.operatorTrivy.additionalVulnerabilityReportFields }},{{ . | lower}}=~\"${{ . | lower}}\"{{- end }}})",
          "instant": true,
          "legendFormat": "Critical",
          "range": false,
          "refId": "A"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${ds_prometheus}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "sum(trivy_vulnerability_id{severity=\"High\", namespace=~\"$namespace\",vuln_id=~\"$vuln_id\",resource_name=~\"$resource_name\",resource=~\"$resource\"{{- range .Values.operatorTrivy.additionalVulnerabilityReportFields }},{{ . | lower}}=~\"${{ . | lower}}\"{{- end }}})",
          "hide": false,
          "instant": true,
          "legendFormat": "High",
          "range": false,
          "refId": "B"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${ds_prometheus}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "sum(trivy_vulnerability_id{severity=\"Medium\", namespace=~\"$namespace\",vuln_id=~\"$vuln_id\",resource_name=~\"$resource_name\",resource=~\"$resource\"{{- range .Values.operatorTrivy.additionalVulnerabilityReportFields }},{{ . | lower}}=~\"${{ . | lower}}\"{{- end }}})",
          "hide": false,
          "instant": true,
          "legendFormat": "Medium",
          "range": false,
          "refId": "C"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${ds_prometheus}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "sum(trivy_vulnerability_id{severity=\"Low\", namespace=~\"$namespace\",vuln_id=~\"$vuln_id\",resource_name=~\"$resource_name\",resource=~\"$resource\"{{- range .Values.operatorTrivy.additionalVulnerabilityReportFields }},{{ . | lower}}=~\"${{ . | lower}}\"{{- end }}})",
          "hide": false,
          "instant": true,
          "legendFormat": "Low",
          "range": false,
          "refId": "D"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${ds_prometheus}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "sum(trivy_vulnerability_id{severity=\"Unknown\", namespace=~\"$namespace\",vuln_id=~\"$vuln_id\",resource_name=~\"$resource_name\",resource=~\"$resource\"{{- range .Values.operatorTrivy.additionalVulnerabilityReportFields }},{{ . | lower}}=~\"${{ . | lower}}\"{{- end }}})",
          "hide": false,
          "instant": true,
          "legendFormat": "Unknown",
          "range": false,
          "refId": "E"
        }
      ],
      "title": "Severity Chart by Image",
      "type": "stat"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "${ds_prometheus}"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "fixed"
          },
          "custom": {
            "align": "left",
            "displayMode": "color-background",
            "filterable": true,
            "inspect": false
          },
          "links": [
            {
              "targetBlank": true,
              "title": "More details",
              "url": "${__data.fields.vuln_url}"
            }
          ],
          "mappings": [
            {
              "options": {
                "pattern": "Critical",
                "result": {
                  "color": "red",
                  "index": 0
                }
              },
              "type": "regex"
            },
            {
              "options": {
                "pattern": "High",
                "result": {
                  "color": "orange",
                  "index": 1
                }
              },
              "type": "regex"
            },
            {
              "options": {
                "pattern": "Medium",
                "result": {
                  "color": "yellow",
                  "index": 2
                }
              },
              "type": "regex"
            },
            {
              "options": {
                "pattern": "Low",
                "result": {
                  "color": "green",
                  "index": 3
                }
              },
              "type": "regex"
            }
          ],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              }
            ]
          }
        },
        "overrides": [
          {
            "matcher": {
              "id": "byName",
              "options": "vuln_url"
            },
            "properties": [
              {
                "id": "custom.hidden",
                "value": true
              }
            ]
          },
          {
            "matcher": {
              "id": "byName",
              "options": "resource"
            },
            "properties": [
              {
                "id": "displayName",
                "value": "Package name"
              }
            ]
          }
        ]
      },
      "gridPos": {
        "h": 14,
        "w": 24,
        "x": 0,
        "y": 9
      },
      "id": 4,
      "options": {
        "footer": {
          "enablePagination": true,
          "fields": "",
          "reducer": [
            "sum"
          ],
          "show": false
        },
        "frameIndex": 1,
        "showHeader": true,
        "sortBy": [
          {
            "desc": false,
            "displayName": "severity"
          }
        ]
      },
      "pluginVersion": "8.5.13",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "${ds_prometheus}"
          },
          "editorMode": "code",
          "exemplar": false,
          "expr": "sum(trivy_vulnerability_id{namespace=~\"$namespace\",severity=~\"$severity\",vuln_id=~\"$vuln_id\",resource_name=~\"$resource_name\",resource=~\"$resource\"{{- range .Values.operatorTrivy.additionalVulnerabilityReportFields }},{{ . | lower}}=~\"${{ . | lower}}\"{{- end }}}) by (image_repository,image_tag,severity, vuln_id, installed_version, fixed_version, namespace, resource_kind, resource, resource_name,vuln_title,vuln_url{{- range .Values.operatorTrivy.additionalVulnerabilityReportFields }},{{ . | lower}}{{- end }})",
          "format": "table",
          "hide": false,
          "instant": true,
          "range": false,
          "refId": "A"
        }
      ],
      "title": "Vulnerability by Vulnerability Id",
      "transformations": [
        {
          "id": "organize",
          "options": {
            "excludeByName": {
              "Time": true,
              "Value": true,
              "vuln_id_hidden": false
            },
            "indexByName": {
              "Time": 0,
              "Value": 8,
              "image_repository": 1,
              "namespace": 2,
              "resource": 3,
              "resource_kind": 9,
              "resource_name": 4,
              "image_tag": 10,
              "installed_version": 20,
              "fixed_version": 21,
              "severity": 5,
              "vuln_id": 6,
              "vuln_title": 7
            },
            "renameByName": {
              "image_repository": ""
            }
          }
        }
      ],
      "type": "table"
    }
  ],
  "schemaVersion": 36,
  "style": "dark",
  "tags": [],
  "templating": {
    "list": [
      {
        "current": {
          "selected": false,
          "text": "main",
          "value": "main"
        },
        "hide": 0,
        "includeAll": false,
        "label": "Prometheus",
        "multi": false,
        "name": "ds_prometheus",
        "options": [],
        "query": "prometheus",
        "refresh": 1,
        "regex": "",
        "skipUrlSync": false,
        "type": "datasource"
      },
      {
        "allValue": ".*",
        "current": {
          "selected": true,
          "text": [
            "All"
          ],
          "value": [
            "$__all"
          ]
        },
        "datasource": {
          "type": "prometheus",
          "uid": "${ds_prometheus}"
        },
        "definition": "label_values(trivy_vulnerability_id, namespace)",
        "hide": 0,
        "includeAll": true,
        "label": "Namespace",
        "multi": true,
        "name": "namespace",
        "options": [],
        "query": {
          "query": "label_values(trivy_vulnerability_id, namespace)",
          "refId": "StandardVariableQuery"
        },
        "refresh": 2,
        "regex": "",
        "skipUrlSync": false,
        "sort": 1,
        "type": "query"
      },
      {
        "allValue": ".*",
        "current": {
          "selected": true,
          "text": [
            "All"
          ],
          "value": [
            "$__all"
          ]
        },
        "datasource": {
          "type": "prometheus",
          "uid": "${ds_prometheus}"
        },
        "definition": "label_values(trivy_vulnerability_id, resource_name)",
        "hide": 0,
        "includeAll": true,
        "label": "Resource name",
        "multi": true,
        "name": "resource_name",
        "options": [],
        "query": {
          "query": "label_values(trivy_vulnerability_id, resource_name)",
          "refId": "StandardVariableQuery"
        },
        "refresh": 2,
        "regex": "",
        "skipUrlSync": false,
        "sort": 1,
        "type": "query"
      },
      {
        "allValue": ".*",
        "current": {
          "selected": true,
          "text": [
            "All"
          ],
          "value": [
            "$__all"
          ]
        },
        "datasource": {
          "type": "prometheus",
          "uid": "${ds_prometheus}"
        },
        "definition": "label_values(trivy_vulnerability_id, severity)",
        "description": "",
        "hide": 0,
        "includeAll": true,
        "label": "Severity",
        "multi": true,
        "name": "severity",
        "options": [],
        "query": {
          "query": "label_values(trivy_vulnerability_id, severity)",
          "refId": "StandardVariableQuery"
        },
        "refresh": 2,
        "regex": "",
        "skipUrlSync": false,
        "sort": 1,
        "type": "query"
      },
      {
        "allValue": ".*",
        "current": {
          "selected": true,
          "text": [
            "All"
          ],
          "value": [
            "$__all"
          ]
        },
        "datasource": {
          "type": "prometheus",
          "uid": "${ds_prometheus}"
        },
        "definition": "label_values(trivy_vulnerability_id, vuln_id)",
        "hide": 0,
        "includeAll": true,
        "label": "Vulnerability Id",
        "multi": true,
        "name": "vuln_id",
        "options": [],
        "query": {
          "query": "label_values(trivy_vulnerability_id, vuln_id)",
          "refId": "StandardVariableQuery"
        },
        "refresh": 2,
        "regex": "",
        "skipUrlSync": false,
        "sort": 1,
        "type": "query"
      },
{{- range .Values.operatorTrivy.additionalVulnerabilityReportFields }}
      {
        "allValue": ".*",
        "current": {
          "selected": true,
          "text": [
            "All"
          ],
          "value": [
            "$__all"
          ]
        },
        "datasource": {
          "type": "prometheus",
          "uid": "${ds_prometheus}"
        },
        "definition": "label_values(trivy_vulnerability_id, {{ . | lower }})",
        "hide": 0,
        "includeAll": true,
        "label": {{ . | quote }},
        "multi": true,
        "name": {{ . | lower | quote }},
        "options": [],
        "query": {
          "query": "label_values(trivy_vulnerability_id, {{ . | lower }})",
          "refId": "StandardVariableQuery"
        },
        "refresh": 2,
        "regex": "",
        "skipUrlSync": false,
        "sort": 1,
        "type": "query"
      },
{{- end }}
      {
        "current": {
          "selected": false,
          "text": [
            "All"
          ],
          "value": [
            "$__all"
          ]
        },
        "datasource": {
          "type": "prometheus",
          "uid": "${ds_prometheus}"
        },
        "definition": "label_values(trivy_vulnerability_id, resource)",
        "hide": 0,
        "includeAll": true,
        "label": "Package name",
        "multi": true,
        "name": "resource",
        "options": [],
        "query": {
          "query": "label_values(trivy_vulnerability_id, resource)",
          "refId": "StandardVariableQuery"
        },
        "refresh": 2,
        "regex": "",
        "skipUrlSync": false,
        "sort": 1,
        "type": "query"
      },
      {
        "datasource": {
          "type": "prometheus",
          "uid": "${ds_prometheus}"
        },
        "filters": [],
        "hide": 0,
        "name": "Filters",
        "skipUrlSync": false,
        "type": "adhoc"
      }
    ]
  },
  "time": {
    "from": "now-30m",
    "to": "now"
  },
  "timepicker": {},
  "timezone": "",
  "title": "Trivy Image Vulnerability Overview",
  "uid": "4SECJjm4z",
  "version": 4,
  "weekStart": ""
}
