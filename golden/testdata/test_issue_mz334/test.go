package testissuemz334

import "github.com/osscontainertools/kaniko/golden/types"

var Tests = types.GoldenTests{
	Name:       "test_issue_mz334",
	Dockerfile: "Dockerfile",
	Tests: []types.GoldenTest{
		{
			Args:       []string{"--no-push", "--cache", "--cache-copy-layers"},
			CachedKeys: []string{},
			Plan:       "plan",
		},
		{
			Args: []string{"--no-push", "--cache", "--cache-copy-layers"},
			Env: map[string]string{
				"FF_KANIKO_CACHE_LOOKAHEAD": "1",
			},
			CachedKeys: []string{
				"e6201ef5c7fab656a2bb0615888d02f181fc7677dbeaaace249484bbdbf79106",
				"09e840ef4fefde5c634154ead19bbff9cfa881073b77cb3c85a400fe215bd518",
				"26d0935bb7ffd5e1f69ecc89608706dfda254c8e9e217f8ccf171d6f3078e0d1",
			},
			Plan: "cached",
		},
		{
			Args: []string{"--no-push", "--cache", "--cache-copy-layers"},
			Env: map[string]string{
				"FF_KANIKO_CACHE_LOOKAHEAD":             "1",
				"FF_KANIKO_INFER_CROSS_STAGE_CACHE_KEY": "1",
				"FF_KANIKO_ROLLING_CACHE_KEY":           "1",
			},
			CachedKeys: []string{
				"f7235dd2645ceb444a287920fe39080370decf958d6463f7ab975033cd634d4e",
				"b82612d288c7f83729db7dc3249f9409dc7029e77c18501792c4907bec2f583f",
				"b4c8d2b4d19b1e9f5db02b9fc802970bd6356c4834df30194611065e85964466",
				"b5547e85c432f089558e3f2f8f9a0669ab164a9cec4f86a79c9aff7ef7bf8657",
				"7f5e1ac9d777aa4407ab852e4d7e43b6f520b6081d00328f9d34824f953d6339",
				"f84d4a33203afe6d505bd43d4e69837b4ed20d3b982a8a18253b483588fd7843",
			},
			Plan: "inferred",
		},
	},
}
