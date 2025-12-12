#!/bin/bash

# Auto-detect lp_solve installation
LP_SOLVE_PREFIX=$(brew --prefix lp_solve 2>/dev/null || echo "/usr/local")

export CGO_CFLAGS="-I${LP_SOLVE_PREFIX}/include"
export CGO_LDFLAGS="-L${LP_SOLVE_PREFIX}/lib -llpsolve55"

go run main.go
