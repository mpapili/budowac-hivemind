package brain

import "math"

// SurfaceHeight mirrors client + server heightmap (seed 42 local-dev).
func SurfaceHeight(wx, wz int, seed int64) int {
	n := fbm(float64(wx)*0.04, float64(wz)*0.04, seed)
	return int(math.Floor(6 + n*14))
}

func hash2(x, z int, seed int64) float64 {
	// Match JS >>> (logical shift) from client generator.ts — not Go >> .
	n := int32(seed) ^ imul(int32(x), 374761393) ^ imul(int32(z), 668265263)
	u := uint32(n)
	n = imul(int32(u^(u>>13)), 1274126177)
	u = uint32(n)
	return float64(u^(u>>16)) / 4294967296.0
}

func imul(a, b int32) int32 { return int32(int64(a) * int64(b)) }

func smoothNoise(x, z float64, seed int64) float64 {
	x0 := int(math.Floor(x))
	z0 := int(math.Floor(z))
	fx := x - float64(x0)
	fz := z - float64(z0)
	sx := fx * fx * (3 - 2*fx)
	sz := fz * fz * (3 - 2*fz)
	n00 := hash2(x0, z0, seed)
	n10 := hash2(x0+1, z0, seed)
	n01 := hash2(x0, z0+1, seed)
	n11 := hash2(x0+1, z0+1, seed)
	ix0 := n00 + (n10-n00)*sx
	ix1 := n01 + (n11-n01)*sx
	return ix0 + (ix1-ix0)*sz
}

func fbm(x, z float64, seed int64) float64 {
	var v, amp, norm float64
	amp = 1
	freq := 1.0
	for i := 0; i < 4; i++ {
		v += smoothNoise(x*freq, z*freq, seed+int64(i*1013)) * amp
		norm += amp
		amp *= 0.5
		freq *= 2
	}
	return v / norm
}
