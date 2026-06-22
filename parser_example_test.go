package audit

import (
	"encoding/json"
	"fmt"

	"github.com/pucora/lura/v2/config"
)

func ExampleParse() {
	cfg, err := config.NewParser().Parse("./tests/example1.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cfg.Normalize()

	result := Parse(&cfg)
	r, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(r))

	// output:
	// {
	//   "d": [
	//     7220,
	//     2000
	//   ],
	//   "a": null,
	//   "e": [
	//     {
	//       "d": [
	//         2,
	//         0,
	//         0,
	//         140000,
	//         0,
	//         0
	//       ],
	//       "b": [
	//         {
	//           "d": [
	//             64
	//           ],
	//           "c": {
	//             "github.com/pucora/pucora-httpcache": [
	//               0
	//             ],
	//             "github.com/pucora/pucora-lua/proxy/backend": [
	//               2
	//             ]
	//           }
	//         }
	//       ],
	//       "c": {
	//         "github.com/pucora/pucora-jose/validator": [],
	//         "github.com/pucora/pucora-lua/proxy": [
	//           3
	//         ],
	//         "modifier/response-body": [
	//           5,
	//           2,
	//           0,
	//           1,
	//           1,
	//           1
	//         ],
	//         "validation/response-json-schema": [
	//           18,
	//           1,
	//           400,
	//           1
	//         ]
	//       }
	//     },
	//     {
	//       "d": [
	//         2,
	//         1,
	//         1,
	//         10000,
	//         7,
	//         0
	//       ],
	//       "b": [
	//         {
	//           "d": [
	//             64
	//           ],
	//           "c": {
	//             "backend/http/client": [
	//               3
	//             ]
	//           }
	//         }
	//       ],
	//       "c": {
	//         "github.com/pucora/pucora/transport/http/client/executor": [
	//           1
	//         ]
	//       }
	//     },
	//     {
	//       "d": [
	//         2,
	//         0,
	//         0,
	//         2000,
	//         0,
	//         0
	//       ],
	//       "b": [
	//         {
	//           "d": [
	//             64
	//           ],
	//           "c": {}
	//         }
	//       ],
	//       "c": {
	//         "websocket": [
	//           27,
	//           4096,
	//           4096,
	//           4096,
	//           3200000,
	//           0,
	//           10000,
	//           60000,
	//           54000,
	//           300000,
	//           1
	//         ]
	//       }
	//     },
	//     {
	//       "d": [
	//         2,
	//         0,
	//         0,
	//         2000,
	//         0,
	//         1
	//       ],
	//       "b": [
	//         {
	//           "d": [
	//             32
	//           ],
	//           "c": {
	//             "backend/soap": [
	//               10,
	//               1,
	//               0,
	//               1
	//             ]
	//           }
	//         }
	//       ],
	//       "c": {}
	//     },
	//     {
	//       "d": [
	//         2,
	//         0,
	//         0,
	//         2000,
	//         0,
	//         0
	//       ],
	//       "b": [
	//         {
	//           "d": [
	//             64
	//           ],
	//           "c": {
	//             "github.com/pucora/pucora-httpcache": [
	//               7
	//             ]
	//           }
	//         }
	//       ],
	//       "c": {}
	//     },
	//     {
	//       "d": [
	//         2,
	//         0,
	//         0,
	//         2000,
	//         0,
	//         1
	//       ],
	//       "b": [
	//         {
	//           "d": [
	//             64
	//           ],
	//           "c": {
	//             "ai/llm": [
	//               1,
	//               0,
	//               0,
	//               1
	//             ]
	//           }
	//         }
	//       ],
	//       "c": {}
	//     },
	//     {
	//       "d": [
	//         2,
	//         0,
	//         0,
	//         2000,
	//         0,
	//         1
	//       ],
	//       "b": [
	//         {
	//           "d": [
	//             64
	//           ],
	//           "c": {
	//             "ai/llm": [
	//               2,
	//               1,
	//               0,
	//               1
	//             ]
	//           }
	//         }
	//       ],
	//       "c": {}
	//     },
	//     {
	//       "d": [
	//         2,
	//         0,
	//         0,
	//         2000,
	//         0,
	//         1
	//       ],
	//       "b": [
	//         {
	//           "d": [
	//             64
	//           ],
	//           "c": {
	//             "ai/llm": [
	//               4,
	//               0,
	//               1,
	//               1
	//             ]
	//           }
	//         }
	//       ],
	//       "c": {}
	//     },
	//     {
	//       "d": [
	//         2,
	//         0,
	//         0,
	//         2000,
	//         0,
	//         1
	//       ],
	//       "b": [
	//         {
	//           "d": [
	//             64
	//           ],
	//           "c": {
	//             "ai/llm": [
	//               8,
	//               1,
	//               1,
	//               1
	//             ]
	//           }
	//         }
	//       ],
	//       "c": {}
	//     },
	//     {
	//       "d": [
	//         2,
	//         0,
	//         0,
	//         2000,
	//         0,
	//         1
	//       ],
	//       "b": [
	//         {
	//           "d": [
	//             64
	//           ],
	//           "c": {}
	//         }
	//       ],
	//       "c": {
	//         "ai/mcp": []
	//       }
	//     },
	//     {
	//       "d": [
	//         2,
	//         0,
	//         0,
	//         2000,
	//         0,
	//         0
	//       ],
	//       "b": [
	//         {
	//           "d": [
	//             64
	//           ],
	//           "c": {}
	//         }
	//       ],
	//       "c": {
	//         "ai/mcp": []
	//       }
	//     },
	//     {
	//       "d": [
	//         2,
	//         0,
	//         0,
	//         2000,
	//         0,
	//         1
	//       ],
	//       "b": [
	//         {
	//           "d": [
	//             64
	//           ],
	//           "c": {}
	//         }
	//       ],
	//       "c": {
	//         "ai/mcp": []
	//       }
	//     },
	//     {
	//       "d": [
	//         2,
	//         0,
	//         0,
	//         2000,
	//         0,
	//         1
	//       ],
	//       "b": [
	//         {
	//           "d": [
	//             64
	//           ],
	//           "c": {}
	//         }
	//       ],
	//       "c": {
	//         "ai/mcp": []
	//       }
	//     },
	//     {
	//       "d": [
	//         2,
	//         0,
	//         0,
	//         10000,
	//         8,
	//         2
	//       ],
	//       "b": [
	//         {
	//           "d": [
	//             64
	//           ],
	//           "c": {}
	//         },
	//         {
	//           "d": [
	//             64
	//           ],
	//           "c": {}
	//         },
	//         {
	//           "d": [
	//             64
	//           ],
	//           "c": {}
	//         }
	//       ],
	//       "c": {
	//         "github.com/pucora/pucora/proxy": [
	//           1
	//         ]
	//       }
	//     }
	//   ],
	//   "c": {
	//     "ai/mcp": [
	//       2,
	//       3
	//     ],
	//     "auth/api-keys": [],
	//     "github.com/pucora/pucora-lua/router": [
	//       1
	//     ],
	//     "github_com/devopsfaith/bloomfilter": [
	//       1,
	//       2,
	//       0
	//     ],
	//     "github_com/luraproject/lura/router/gin": [],
	//     "github_com/pucora/pucora-influx": [],
	//     "github_com/pucora/pucora/transport/http/server/handler": [
	//       4
	//     ],
	//     "grpc": [
	//       1
	//     ],
	//     "modifier/response-headers": [
	//       15
	//     ],
	//     "qos/ratelimit/service": [],
	//     "telemetry/opentelemetry": [
	//       50,
	//       100,
	//       1,
	//       2,
	//       1
	//     ]
	//   },
	//   "ap": {}
	// }
}

func ExampleParse_withRevokerServer() {
	cfg, err := config.NewParser().Parse("./tests/revoker.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cfg.Normalize()

	result := Parse(&cfg)
	r, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(r))

	// output:
	// {
	//   "d": [
	//     0,
	//     2000
	//   ],
	//   "a": null,
	//   "e": [
	//     {
	//       "d": [
	//         2,
	//         0,
	//         0,
	//         2000,
	//         0,
	//         0
	//       ],
	//       "b": [
	//         {
	//           "d": [
	//             64
	//           ],
	//           "c": {}
	//         }
	//       ],
	//       "c": {}
	//     }
	//   ],
	//   "c": {
	//     "github_com/devopsfaith/bloomfilter": [
	//       1,
	//       1,
	//       1
	//     ]
	//   },
	//   "ap": {}
	// }
}
