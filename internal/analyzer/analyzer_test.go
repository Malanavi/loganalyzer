package analyzer

import "testing"

func TestStats(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  StatsResult
	}{
		{
			name:  "empty log",
			lines: []string{},
			want:  StatsResult{},
		},
		{
			name: "info only",
			lines: []string{
				"[INFO] started",
				"[INFO] running",
			},
			want: StatsResult{
				Total: 2,
				Info:  2,
			},
		},
		{
			name: "mixed levels",
			lines: []string{
				"[INFO] started",
				"[WARN] slow",
				"[ERROR] failed",
				"random line",
			},
			want: StatsResult{
				Total: 4,
				Info:  1,
				Warn:  1,
				Error: 1,
			},
		},
		{
			name: "unknown prefix",
			lines: []string{
				"[DEBUG] something",
			},
			want: StatsResult{
				Total: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Stats(tt.lines)
			if got != tt.want {
				t.Errorf("Stats() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestErrors(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  []string
	}{
		{
			name:  "no errors",
			lines: []string{"[INFO] ok"},
			want:  []string{},
		},
		{
			name: "multiple errors",
			lines: []string{
				"[ERROR] first",
				"[INFO] ok",
				"[ERROR] second",
			},
			want: []string{
				"[ERROR] first",
				"[ERROR] second",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Errors(tt.lines)
			if len(got) != len(tt.want) {
				t.Fatalf("expected %d errors, got %d", len(tt.want), len(got))
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("error[%d] = %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestTop(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  []TopLogItem
	}{
		{
			name:  "empty log",
			lines: []string{},
			want:  []TopLogItem{},
		},
		{
			name: "sorted by frequency",
			lines: []string{
				"[INFO] ok",
				"[WARN] slow",
				"[INFO] ok",
				"[INFO] ok",
				"[WARN] slow",
			},
			want: []TopLogItem{
				{Log: "[INFO] ok", Times: 3},
				{Log: "[WARN] slow", Times: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Top(tt.lines)
			if len(got) != len(tt.want) {
				t.Fatalf("expected %d items, got %d", len(tt.want), len(got))
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("top[%d] = %+v, want %+v", i, got[i], tt.want[i])
				}
			}
		})
	}
}
