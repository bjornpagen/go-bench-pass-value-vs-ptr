package main

import "testing"

func Benchmark1B(b *testing.B) {
	var p payload1B
	b.Run("passByPtr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			B1PassByPtr(&p)
		}
	})
	b.Run("passByValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			B1PassByVal(p)
		}
	})
}

func Benchmark4B(b *testing.B) {
	var p payload4B
	b.Run("passByPtr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			B4PassByPtr(&p)
		}
	})
	b.Run("passByValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			B4PassByVal(p)
		}
	})
}

func Benchmark16B(b *testing.B) {
	var p payload16B
	b.Run("passByPtr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			B16PassByPtr(&p)
		}
	})
	b.Run("passByValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			B16PassByVal(p)
		}
	})
}

func Benchmark64B(b *testing.B) {
	var p payload64B
	b.Run("passByPtr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			B64PassByPtr(&p)
		}
	})
	b.Run("passByValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			B64PassByVal(p)
		}
	})
}

func Benchmark128B(b *testing.B) {
	var p payload128B
	b.Run("passByPtr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			B128PassByPtr(&p)
		}
	})
	b.Run("passByValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			B128PassByVal(p)
		}
	})
}

func Benchmark256B(b *testing.B) {
	var p payload256B
	b.Run("passByPtr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			B256PassByPtr(&p)
		}
	})
	b.Run("passByValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			B256PassByVal(p)
		}
	})
}

func Benchmark1K(b *testing.B) {
	var p payload1KiB
	b.Run("passByPtr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			K1PassByPtr(&p)
		}
	})
	b.Run("passByValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			K1PassByVal(p)
		}
	})
}

func Benchmark4K(b *testing.B) {
	var p payload4KiB
	b.Run("passByPtr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			K4PassByPtr(&p)
		}
	})
	b.Run("passByValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			K4PassByVal(p)
		}
	})
}

func Benchmark16K(b *testing.B) {
	var p payload16KiB
	b.Run("passByPtr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			K16PassByPtr(&p)
		}
	})
	b.Run("passByValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			K16PassByVal(p)
		}
	})
}

func Benchmark64K(b *testing.B) {
	var p payload64KiB
	b.Run("passByPtr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			K64PassByPtr(&p)
		}
	})
	b.Run("passByValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			K64PassByVal(p)
		}
	})
}

func Benchmark256K(b *testing.B) {
	var p payload256KiB
	b.Run("passByPtr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			K256PassByPtr(&p)
		}
	})
	b.Run("passByValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			K256PassByVal(p)
		}
	})
}

func Benchmark1M(b *testing.B) {
	var p payload1MiB
	b.Run("passByPtr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			M1PassByPtr(&p)
		}
	})
	b.Run("passByValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			M1PassByVal(p)
		}
	})
}
