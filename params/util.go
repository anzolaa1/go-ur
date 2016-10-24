// Copyright 2015 The go-ur Authors
// This file is part of the go-ur library.
//
// The go-ur library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ur library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ur library. If not, see <http://www.gnu.org/licenses/>.

package params

import "math/big"

var (
	TestNetHomesteadBlock           = big.NewInt(1)       // Testnet homestead block
	MainNetHomesteadBlock           = big.NewInt(1)       // Mainnet homestead block
	TestNetHomesteadGasRepriceBlock = big.NewInt(1783000) // Test net gas reprice block
	MainNetHomesteadGasRepriceBlock = big.NewInt(2463000) // Main net gas reprice block
)
