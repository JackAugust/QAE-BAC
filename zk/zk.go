package zk

import (
	"crypto/rand"
	"math/big"
)

// a1 is 要求的最低救治率；a2 is 急救中心的救治率
func ZK(a1, a2 int64) bool {
	m := big.NewInt(a2)
	n := big.NewInt(a1)
	// fmt.Println("m is ", m, " and n is ", n)
	p := big.NewInt(47) // 选取一个大质数p
	g := big.NewInt(7)  // 选择一个大于1且小于p的整数g
	p1 := new(big.Int).Sub(p, big.NewInt(1))
	// 1. 证明者A随机生成一个数据 作为私钥
	// r := big.NewInt(2) // 选择一个大于1且小于p的整数g
	r, _ := rand.Int(rand.Reader, p1)
	r.Add(r, big.NewInt(1))
	// 证明者A公钥 z = g^r mod p {p,g,z}={7,2,4}
	// z := new(big.Int).Exp(g, r, p)

	// 2. 验证者V选择一个大于1且小于p随机值作为挑战值b
	b, _ := rand.Int(rand.Reader, p1)
	b.Add(b, big.NewInt(1))

	// 3. 证明者A计算x = b * g^(m-n+1)r mod p 和y = b * g^(m+1) mod p
	k := new(big.Int).Sub(m, n)
	k.Add(k, big.NewInt(1))
	k.Mul(k, r)

	// 得到密文x
	x := new(big.Int)
	q := new(big.Float)
	if k.Cmp(big.NewInt(1)) < 0 {
		k.Sub(big.NewInt(1), k)
		x.Exp(g, k, nil)
		q.SetInt(x)
		q.Quo(big.NewFloat(1.0), q)
		x, _ = q.Int(nil)
	} else {
		x.Exp(g, k, nil)
	}
	x.Mul(x, b)
	x.Mod(x, p)

	// 得到密文y
	k2 := new(big.Int).Add(m, big.NewInt(1))
	k2.Mul(k2, r)
	y := new(big.Int).Exp(g, k2, p)
	y.Mul(y, b)
	y.Mod(y, p)

	// 把x发送给验证着V

	// v := new(big.Int).Exp(g, b, p)

	// 验证者V发送公钥v给证明者A
	// A计算t = b * g^r mod p
	// t := new(big.Int).Set(z)
	k3 := new(big.Int).Mul(n, r)
	s := new(big.Int).Exp(g, k3, p)
	s.Mul(s, x)
	s.Mod(s, p)

	// 验证者V检查y是否等于 x * n^b mod p
	// 如果是，则认为证明有效，即m > n
	// y := new(big.Int).Set(x)
	// y.Add(y, b)
	// fmt.Println("x = ", x, " z = ", z, "y = ", y)
	// nb = x * n^b mod p
	// nb := new(big.Int).Exp(n, r, p)
	// nb.Mul(y, nb)
	// nb.Mod(nb, p)
	// fmt.Println("nb = ", nb)
	// fmt.Println("x = ", x, " s = ", s, "y = ", y, " (m-n+1)*r = ", k, "and b = ", b, " and r = ", r, " and k3 = ", k3, " and k2 = ", k2)

	// x=(m-n+b)*g^r mod p
	if s.Cmp(y) == 0 {
		// fmt.Println("有效的证明: ", s.Cmp(y))
		return true
	} else {
		// fmt.Println("无效的证明: ", s.Cmp(y))
		return false
	}
}
