package comparekv

import "testing"

func TestCompareFields(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareFields(tt.args.struct1, tt.args.struct2, tt.args.fields...); got != tt.want {
				t.Errorf("CompareFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCompareFields(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			CompareFields(tt.args.struct1, tt.args.struct2, tt.args.fields...)
		}
	}
}
