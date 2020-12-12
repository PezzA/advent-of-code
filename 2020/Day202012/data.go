package Day202012

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2020, 12, "Rain Risk"
}

func (td dayEntry) PuzzleInput() string {
	return `W1
L90
F26
S3
W2
N5
L180
S4
F41
W1
F48
W3
F44
F63
W5
N3
E5
F7
R180
W1
N3
W3
R180
F38
N1
E3
L90
F49
S5
F11
E1
R90
W3
N4
E3
R180
W5
F93
S1
F67
L180
W3
S5
N3
L180
E1
N2
F90
S3
W3
N3
E1
F77
W3
N3
L180
E2
F25
N1
W1
S1
E5
S4
R180
W1
F13
L90
W4
S5
F13
N2
R90
F89
R90
W3
L90
W5
N5
E4
S4
F26
N1
R270
N5
E5
L180
R180
W4
R90
W5
F49
S2
F53
W5
L180
F54
L90
N3
F3
W5
R180
S4
L90
F49
W4
S5
F73
L270
W1
S4
F46
S4
W2
S5
E4
N4
F6
E4
S4
F38
F4
E4
R90
N1
W2
F20
N3
F64
R90
S3
W5
N1
W4
N3
F17
R90
F62
E4
E2
F47
R90
W1
N3
R90
N2
R90
F57
L270
N2
W3
S5
W4
E2
S5
F93
F57
S1
L90
F50
N2
F4
N1
L90
F34
N2
L90
N4
W3
N5
R180
S3
L90
W3
S3
L90
F70
L90
E1
F92
N1
F96
F85
S5
L90
W1
L90
W1
F23
L90
S1
R90
W5
S5
F66
W3
L180
W2
L90
N2
E3
R270
R270
N3
W5
R90
S3
E1
R90
F78
E1
S1
R90
S3
F52
S4
F9
L90
W1
N2
F8
R90
N1
F63
E5
F18
E3
F43
E2
F10
R90
F96
S5
F22
W2
S5
F39
R90
F38
S5
R90
E3
L90
W3
N2
F14
L270
S4
F78
F85
L90
N3
E3
S3
F98
E2
S2
F100
S3
S3
W5
W3
S5
F67
L180
S2
E5
S1
L90
N5
E2
W2
R90
E1
N2
L90
F77
W1
F84
L90
S2
E4
R90
E1
R90
S3
S4
F89
R90
N1
E4
R90
N1
F97
L90
S1
W3
R180
F70
S5
E1
L180
W5
F86
S3
F20
R90
S1
W4
R90
W1
F3
S3
R90
F43
L180
F81
E2
N3
F16
L90
S2
F17
E3
F1
E4
F17
W3
N3
W5
S3
W4
F60
E3
E1
S5
L90
E2
S5
F19
E2
R90
F20
R180
S4
F9
R90
N5
W5
F56
N2
L180
N1
E5
L90
F15
W4
F26
R90
W2
F19
S3
W1
R90
W5
R180
W4
N2
F86
N5
E3
W3
N3
L270
W3
F42
N5
W2
R180
W2
R180
S4
R90
F55
S3
R90
S3
E3
R90
F11
S4
F38
W1
L90
F8
R90
E5
R90
W1
W5
S2
F2
F92
S3
F77
S5
R90
F24
E3
R90
N3
F16
L270
W3
F83
L270
E2
F98
L180
F89
E5
F98
S4
E2
L90
N4
L180
F57
S5
R90
L90
S4
W4
S5
S4
W4
F43
N2
F29
W3
R90
F41
R90
N2
F78
R90
E5
N1
W2
F6
L270
W5
F91
W5
N1
S4
F41
W4
F74
E1
R90
N4
F76
W4
S2
L180
N2
R180
W4
F79
R270
W1
F92
W1
L90
F71
N4
L180
W4
F16
W5
F84
S5
F35
W4
R90
F25
L180
N1
E3
F15
S4
R180
F46
S1
W1
R180
E4
N5
R90
S1
W3
S3
L270
F94
R180
N1
W4
N5
W2
S2
W3
F53
L180
S3
F19
N3
F54
L180
S5
F8
S1
N5
L90
E4
N3
F28
R180
F23
E1
L90
E3
F6
W4
R90
N1
F89
S1
W2
S5
F8
N3
F23
N4
F5
L90
N3
R90
W4
L180
S3
F7
N2
W3
R180
E1
L180
S4
R90
S1
F99
N3
F96
W3
R90
F73
W5
F71
R180
S2
F84
N4
F4
W4
R90
F34
E2
W2
F53
N4
R90
N5
E5
R90
F60
N4
F28
S2
W1
N4
F54
R270
F45
S5
F93
L90
F66
R180
F92
N4
F97
R90
W5
S1
W5
F68
S3
L90
E3
F94
S4
F64
R180
F18
N1
S4
E5
E2
F81
N1
L90
F3
R90
F81
W4
S4
E5
N5
R270
E3
S2
W1
L180
S1
F84
W2
L270
F6
N1
R180
E5
F7
E2
L180
E2
F80
N1
L90
F88
R90
W5
N1
F71
R180
N2
E2
R90
N1
W1
L90
E4
S4
L180
F27
L90
F57
F38
W5
L180
S5
R90
S4
W1
S3
L90
F36
S1
W5
R90
F65
R90
E1
N4
E1
F14
L90
F44
R90
F34
S2
L90
R180
F87
W3
L90
F9
R90
W3
L90
S5
F69
S3
W4
N4
F30
W5
F15
R90
L180
W4
F5
R180
E1
F6
R180
S1
F20
E1
S2
E5
F13
N5
F83
W2
L270
E2
R90
S5
F62
R270
N4
R90
F20
L90
N2
E3
L90
F37
N2
N2
F82
L90
F23
E3
F63
R180
F1
N2
R90
F68
E5
F75
R90
W3
R180
E4
E1
N3
R90
N3
L180
F92
R90
S4
F27
R180
S4
L180
W5
F70
S5
L180
F89
R90
W2
N3
F64
L90
E1
L90
F77
E4
F55
E2
L90
W2
F46
N2
R90
F94
S5
R180
F9
L180
S4
L90
F25`
}

func (td dayEntry) PuzzleSpec() string {
	return ``
}
