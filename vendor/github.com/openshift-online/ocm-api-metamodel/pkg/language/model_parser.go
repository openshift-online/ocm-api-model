// Code generated from ModelParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package language // ModelParser
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import (
	"github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
	"github.com/openshift-online/ocm-api-metamodel/pkg/names"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 34, 330,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 7, 2, 74, 10, 2, 12, 2, 14, 2,
	77, 11, 2, 3, 3, 3, 3, 5, 3, 81, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 87,
	10, 4, 3, 5, 7, 5, 90, 10, 5, 12, 5, 14, 5, 93, 11, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 7, 5, 99, 10, 5, 12, 5, 14, 5, 102, 11, 5, 3, 5, 3, 5, 3, 6, 7, 6,
	107, 10, 6, 12, 6, 14, 6, 110, 11, 6, 3, 6, 5, 6, 113, 10, 6, 3, 6, 3,
	6, 3, 7, 7, 7, 118, 10, 7, 12, 7, 14, 7, 121, 11, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 7, 7, 127, 10, 7, 12, 7, 14, 7, 130, 11, 7, 3, 7, 3, 7, 3, 8, 7,
	8, 135, 10, 8, 12, 8, 14, 8, 138, 11, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8,
	144, 10, 8, 12, 8, 14, 8, 147, 11, 8, 3, 8, 3, 8, 3, 9, 7, 9, 152, 10,
	9, 12, 9, 14, 9, 155, 11, 9, 3, 9, 5, 9, 158, 10, 9, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 5, 11, 168, 10, 11, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 7, 15,
	182, 10, 15, 12, 15, 14, 15, 185, 11, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7,
	15, 191, 10, 15, 12, 15, 14, 15, 194, 11, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	5, 16, 200, 10, 16, 3, 17, 7, 17, 203, 10, 17, 12, 17, 14, 17, 206, 11,
	17, 3, 17, 5, 17, 209, 10, 17, 3, 17, 3, 17, 3, 17, 7, 17, 214, 10, 17,
	12, 17, 14, 17, 217, 11, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 7, 19,
	224, 10, 19, 12, 19, 14, 19, 227, 11, 19, 3, 19, 5, 19, 230, 10, 19, 3,
	19, 7, 19, 233, 10, 19, 12, 19, 14, 19, 236, 11, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 5, 19, 242, 10, 19, 3, 20, 3, 20, 5, 20, 246, 10, 20, 3, 21, 7,
	21, 249, 10, 21, 12, 21, 14, 21, 252, 11, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	7, 21, 258, 10, 21, 12, 21, 14, 21, 261, 11, 21, 3, 21, 3, 21, 3, 22, 3,
	22, 5, 22, 267, 10, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 25,
	3, 25, 3, 26, 7, 26, 278, 10, 26, 12, 26, 14, 26, 281, 11, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 7, 26, 287, 10, 26, 12, 26, 14, 26, 290, 11, 26, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 5, 29, 302,
	10, 29, 3, 30, 3, 30, 7, 30, 306, 10, 30, 12, 30, 14, 30, 309, 11, 30,
	3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 5, 32, 320,
	10, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36,
	2, 2, 37, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
	2, 4, 4, 2, 3, 3, 10, 10, 4, 2, 8, 8, 18, 18, 2, 331, 2, 75, 3, 2, 2, 2,
	4, 80, 3, 2, 2, 2, 6, 86, 3, 2, 2, 2, 8, 91, 3, 2, 2, 2, 10, 108, 3, 2,
	2, 2, 12, 119, 3, 2, 2, 2, 14, 136, 3, 2, 2, 2, 16, 153, 3, 2, 2, 2, 18,
	162, 3, 2, 2, 2, 20, 167, 3, 2, 2, 2, 22, 169, 3, 2, 2, 2, 24, 171, 3,
	2, 2, 2, 26, 175, 3, 2, 2, 2, 28, 183, 3, 2, 2, 2, 30, 199, 3, 2, 2, 2,
	32, 204, 3, 2, 2, 2, 34, 220, 3, 2, 2, 2, 36, 225, 3, 2, 2, 2, 38, 245,
	3, 2, 2, 2, 40, 250, 3, 2, 2, 2, 42, 266, 3, 2, 2, 2, 44, 268, 3, 2, 2,
	2, 46, 271, 3, 2, 2, 2, 48, 274, 3, 2, 2, 2, 50, 279, 3, 2, 2, 2, 52, 293,
	3, 2, 2, 2, 54, 295, 3, 2, 2, 2, 56, 298, 3, 2, 2, 2, 58, 303, 3, 2, 2,
	2, 60, 312, 3, 2, 2, 2, 62, 319, 3, 2, 2, 2, 64, 321, 3, 2, 2, 2, 66, 323,
	3, 2, 2, 2, 68, 325, 3, 2, 2, 2, 70, 327, 3, 2, 2, 2, 72, 74, 5, 4, 3,
	2, 73, 72, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76,
	3, 2, 2, 2, 76, 3, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 81, 5, 6, 4, 2,
	79, 81, 5, 28, 15, 2, 80, 78, 3, 2, 2, 2, 80, 79, 3, 2, 2, 2, 81, 5, 3,
	2, 2, 2, 82, 87, 5, 8, 5, 2, 83, 87, 5, 12, 7, 2, 84, 87, 5, 14, 8, 2,
	85, 87, 5, 50, 26, 2, 86, 82, 3, 2, 2, 2, 86, 83, 3, 2, 2, 2, 86, 84, 3,
	2, 2, 2, 86, 85, 3, 2, 2, 2, 87, 7, 3, 2, 2, 2, 88, 90, 5, 56, 29, 2, 89,
	88, 3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2,
	2, 92, 94, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 94, 95, 7, 6, 2, 2, 95, 96,
	5, 70, 36, 2, 96, 100, 7, 22, 2, 2, 97, 99, 5, 10, 6, 2, 98, 97, 3, 2,
	2, 2, 99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101,
	103, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 104, 7, 23, 2, 2, 104, 9, 3,
	2, 2, 2, 105, 107, 5, 56, 29, 2, 106, 105, 3, 2, 2, 2, 107, 110, 3, 2,
	2, 2, 108, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 112, 3, 2, 2, 2,
	110, 108, 3, 2, 2, 2, 111, 113, 7, 19, 2, 2, 112, 111, 3, 2, 2, 2, 112,
	113, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 115, 5, 70, 36, 2, 115, 11,
	3, 2, 2, 2, 116, 118, 5, 56, 29, 2, 117, 116, 3, 2, 2, 2, 118, 121, 3,
	2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 122, 3, 2, 2,
	2, 121, 119, 3, 2, 2, 2, 122, 123, 7, 4, 2, 2, 123, 124, 5, 70, 36, 2,
	124, 128, 7, 22, 2, 2, 125, 127, 5, 16, 9, 2, 126, 125, 3, 2, 2, 2, 127,
	130, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 131,
	3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 131, 132, 7, 23, 2, 2, 132, 13, 3, 2,
	2, 2, 133, 135, 5, 56, 29, 2, 134, 133, 3, 2, 2, 2, 135, 138, 3, 2, 2,
	2, 136, 134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 139, 3, 2, 2, 2, 138,
	136, 3, 2, 2, 2, 139, 140, 7, 16, 2, 2, 140, 141, 5, 70, 36, 2, 141, 145,
	7, 22, 2, 2, 142, 144, 5, 16, 9, 2, 143, 142, 3, 2, 2, 2, 144, 147, 3,
	2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 148, 3, 2, 2,
	2, 147, 145, 3, 2, 2, 2, 148, 149, 7, 23, 2, 2, 149, 15, 3, 2, 2, 2, 150,
	152, 5, 56, 29, 2, 151, 150, 3, 2, 2, 2, 152, 155, 3, 2, 2, 2, 153, 151,
	3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 157, 3, 2, 2, 2, 155, 153, 3, 2,
	2, 2, 156, 158, 5, 18, 10, 2, 157, 156, 3, 2, 2, 2, 157, 158, 3, 2, 2,
	2, 158, 159, 3, 2, 2, 2, 159, 160, 5, 70, 36, 2, 160, 161, 5, 20, 11, 2,
	161, 17, 3, 2, 2, 2, 162, 163, 9, 2, 2, 2, 163, 19, 3, 2, 2, 2, 164, 168,
	5, 22, 12, 2, 165, 168, 5, 24, 13, 2, 166, 168, 5, 26, 14, 2, 167, 164,
	3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 167, 166, 3, 2, 2, 2, 168, 21, 3, 2,
	2, 2, 169, 170, 5, 70, 36, 2, 170, 23, 3, 2, 2, 2, 171, 172, 7, 24, 2,
	2, 172, 173, 7, 25, 2, 2, 173, 174, 5, 70, 36, 2, 174, 25, 3, 2, 2, 2,
	175, 176, 7, 24, 2, 2, 176, 177, 5, 70, 36, 2, 177, 178, 7, 25, 2, 2, 178,
	179, 5, 70, 36, 2, 179, 27, 3, 2, 2, 2, 180, 182, 5, 56, 29, 2, 181, 180,
	3, 2, 2, 2, 182, 185, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 183, 184, 3, 2,
	2, 2, 184, 186, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 186, 187, 7, 15, 2, 2,
	187, 188, 5, 70, 36, 2, 188, 192, 7, 22, 2, 2, 189, 191, 5, 30, 16, 2,
	190, 189, 3, 2, 2, 2, 191, 194, 3, 2, 2, 2, 192, 190, 3, 2, 2, 2, 192,
	193, 3, 2, 2, 2, 193, 195, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 195, 196,
	7, 23, 2, 2, 196, 29, 3, 2, 2, 2, 197, 200, 5, 32, 17, 2, 198, 200, 5,
	40, 21, 2, 199, 197, 3, 2, 2, 2, 199, 198, 3, 2, 2, 2, 200, 31, 3, 2, 2,
	2, 201, 203, 5, 56, 29, 2, 202, 201, 3, 2, 2, 2, 203, 206, 3, 2, 2, 2,
	204, 202, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 208, 3, 2, 2, 2, 206,
	204, 3, 2, 2, 2, 207, 209, 7, 12, 2, 2, 208, 207, 3, 2, 2, 2, 208, 209,
	3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 211, 5, 70, 36, 2, 211, 215, 7,
	22, 2, 2, 212, 214, 5, 34, 18, 2, 213, 212, 3, 2, 2, 2, 214, 217, 3, 2,
	2, 2, 215, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 218, 3, 2, 2, 2,
	217, 215, 3, 2, 2, 2, 218, 219, 7, 23, 2, 2, 219, 33, 3, 2, 2, 2, 220,
	221, 5, 36, 19, 2, 221, 35, 3, 2, 2, 2, 222, 224, 5, 56, 29, 2, 223, 222,
	3, 2, 2, 2, 224, 227, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 225, 226, 3, 2,
	2, 2, 226, 229, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 228, 230, 7, 14, 2, 2,
	229, 228, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 234, 3, 2, 2, 2, 231,
	233, 5, 38, 20, 2, 232, 231, 3, 2, 2, 2, 233, 236, 3, 2, 2, 2, 234, 232,
	3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 237, 3, 2, 2, 2, 236, 234, 3, 2,
	2, 2, 237, 238, 5, 70, 36, 2, 238, 241, 5, 20, 11, 2, 239, 240, 7, 28,
	2, 2, 240, 242, 5, 62, 32, 2, 241, 239, 3, 2, 2, 2, 241, 242, 3, 2, 2,
	2, 242, 37, 3, 2, 2, 2, 243, 246, 7, 9, 2, 2, 244, 246, 7, 13, 2, 2, 245,
	243, 3, 2, 2, 2, 245, 244, 3, 2, 2, 2, 246, 39, 3, 2, 2, 2, 247, 249, 5,
	56, 29, 2, 248, 247, 3, 2, 2, 2, 249, 252, 3, 2, 2, 2, 250, 248, 3, 2,
	2, 2, 250, 251, 3, 2, 2, 2, 251, 253, 3, 2, 2, 2, 252, 250, 3, 2, 2, 2,
	253, 254, 7, 11, 2, 2, 254, 255, 5, 70, 36, 2, 255, 259, 7, 22, 2, 2, 256,
	258, 5, 42, 22, 2, 257, 256, 3, 2, 2, 2, 258, 261, 3, 2, 2, 2, 259, 257,
	3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260, 262, 3, 2, 2, 2, 261, 259, 3, 2,
	2, 2, 262, 263, 7, 23, 2, 2, 263, 41, 3, 2, 2, 2, 264, 267, 5, 44, 23,
	2, 265, 267, 5, 46, 24, 2, 266, 264, 3, 2, 2, 2, 266, 265, 3, 2, 2, 2,
	267, 43, 3, 2, 2, 2, 268, 269, 7, 17, 2, 2, 269, 270, 5, 48, 25, 2, 270,
	45, 3, 2, 2, 2, 271, 272, 7, 20, 2, 2, 272, 273, 5, 70, 36, 2, 273, 47,
	3, 2, 2, 2, 274, 275, 5, 70, 36, 2, 275, 49, 3, 2, 2, 2, 276, 278, 5, 56,
	29, 2, 277, 276, 3, 2, 2, 2, 278, 281, 3, 2, 2, 2, 279, 277, 3, 2, 2, 2,
	279, 280, 3, 2, 2, 2, 280, 282, 3, 2, 2, 2, 281, 279, 3, 2, 2, 2, 282,
	283, 7, 7, 2, 2, 283, 284, 5, 70, 36, 2, 284, 288, 7, 22, 2, 2, 285, 287,
	5, 52, 27, 2, 286, 285, 3, 2, 2, 2, 287, 290, 3, 2, 2, 2, 288, 286, 3,
	2, 2, 2, 288, 289, 3, 2, 2, 2, 289, 291, 3, 2, 2, 2, 290, 288, 3, 2, 2,
	2, 291, 292, 7, 23, 2, 2, 292, 51, 3, 2, 2, 2, 293, 294, 5, 54, 28, 2,
	294, 53, 3, 2, 2, 2, 295, 296, 7, 5, 2, 2, 296, 297, 7, 29, 2, 2, 297,
	55, 3, 2, 2, 2, 298, 299, 7, 21, 2, 2, 299, 301, 5, 70, 36, 2, 300, 302,
	5, 58, 30, 2, 301, 300, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 57, 3, 2,
	2, 2, 303, 307, 7, 26, 2, 2, 304, 306, 5, 60, 31, 2, 305, 304, 3, 2, 2,
	2, 306, 309, 3, 2, 2, 2, 307, 305, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308,
	310, 3, 2, 2, 2, 309, 307, 3, 2, 2, 2, 310, 311, 7, 27, 2, 2, 311, 59,
	3, 2, 2, 2, 312, 313, 5, 70, 36, 2, 313, 314, 7, 28, 2, 2, 314, 315, 5,
	62, 32, 2, 315, 61, 3, 2, 2, 2, 316, 320, 5, 64, 33, 2, 317, 320, 5, 66,
	34, 2, 318, 320, 5, 68, 35, 2, 319, 316, 3, 2, 2, 2, 319, 317, 3, 2, 2,
	2, 319, 318, 3, 2, 2, 2, 320, 63, 3, 2, 2, 2, 321, 322, 9, 3, 2, 2, 322,
	65, 3, 2, 2, 2, 323, 324, 7, 29, 2, 2, 324, 67, 3, 2, 2, 2, 325, 326, 7,
	30, 2, 2, 326, 69, 3, 2, 2, 2, 327, 328, 7, 31, 2, 2, 328, 71, 3, 2, 2,
	2, 35, 75, 80, 86, 91, 100, 108, 112, 119, 128, 136, 145, 153, 157, 167,
	183, 192, 199, 204, 208, 215, 225, 229, 234, 241, 245, 250, 259, 266, 279,
	288, 301, 307, 319,
}
var literalNames = []string{
	"", "'attribute'", "'class'", "'code'", "'enum'", "'error'", "'false'",
	"'in'", "'link'", "'locator'", "'method'", "'out'", "'parameter'", "'resource'",
	"'struct'", "'target'", "'true'", "'value'", "'variable'", "'@'", "'{'",
	"'}'", "'['", "']'", "'('", "')'", "'='",
}
var symbolicNames = []string{
	"", "ATTRIBUTE", "CLASS", "CODE", "ENUM", "ERROR", "FALSE", "IN", "LINK",
	"LOCATOR", "METHOD", "OUT", "PARAMETER", "RESOURCE", "STRUCT", "TARGET",
	"TRUE", "VALUE", "VARIABLE", "AT_SIGN", "LEFT_CURLY_BRACKET", "RIGHT_CURLY_BRACKET",
	"LEFT_SQUARE_BRACKET", "RIGHT_SQUARE_BRACKET", "LEFT_PARENTHESIS", "RIGHT_PARENTHESIS",
	"EQUALS_SIGN", "INTEGER_LITERAL", "STRING_LITERAL", "IDENTIFIER", "LINE_COMMENT",
	"BLOCK_COMMENT", "WS",
}

var ruleNames = []string{
	"file", "declaration", "typeDecl", "enumDecl", "enumMemberDecl", "classDecl",
	"structDecl", "structMemberDecl", "attributeKind", "typeReference", "plainTypeReference",
	"listTypeReference", "mapTypeReference", "resourceDecl", "resourceMemberDecl",
	"methodDecl", "methodMemberDecl", "methodParameterDecl", "parameterDirection",
	"locatorDecl", "locatorMemberDecl", "locatorTargetDecl", "locatorVariableDecl",
	"resourceReference", "errorDecl", "errorMemberDecl", "errorCodeDecl", "annotation",
	"annotationParameters", "annotationParameter", "literal", "booleanLiteral",
	"integerLiteral", "stringLiteral", "identifier",
}

type ModelParser struct {
	*antlr.BaseParser
}

// NewModelParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ModelParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewModelParser(input antlr.TokenStream) *ModelParser {
	this := new(ModelParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ModelParser.g4"

	return this
}

// ModelParser tokens.
const (
	ModelParserEOF                  = antlr.TokenEOF
	ModelParserATTRIBUTE            = 1
	ModelParserCLASS                = 2
	ModelParserCODE                 = 3
	ModelParserENUM                 = 4
	ModelParserERROR                = 5
	ModelParserFALSE                = 6
	ModelParserIN                   = 7
	ModelParserLINK                 = 8
	ModelParserLOCATOR              = 9
	ModelParserMETHOD               = 10
	ModelParserOUT                  = 11
	ModelParserPARAMETER            = 12
	ModelParserRESOURCE             = 13
	ModelParserSTRUCT               = 14
	ModelParserTARGET               = 15
	ModelParserTRUE                 = 16
	ModelParserVALUE                = 17
	ModelParserVARIABLE             = 18
	ModelParserAT_SIGN              = 19
	ModelParserLEFT_CURLY_BRACKET   = 20
	ModelParserRIGHT_CURLY_BRACKET  = 21
	ModelParserLEFT_SQUARE_BRACKET  = 22
	ModelParserRIGHT_SQUARE_BRACKET = 23
	ModelParserLEFT_PARENTHESIS     = 24
	ModelParserRIGHT_PARENTHESIS    = 25
	ModelParserEQUALS_SIGN          = 26
	ModelParserINTEGER_LITERAL      = 27
	ModelParserSTRING_LITERAL       = 28
	ModelParserIDENTIFIER           = 29
	ModelParserLINE_COMMENT         = 30
	ModelParserBLOCK_COMMENT        = 31
	ModelParserWS                   = 32
)

// ModelParser rules.
const (
	ModelParserRULE_file                 = 0
	ModelParserRULE_declaration          = 1
	ModelParserRULE_typeDecl             = 2
	ModelParserRULE_enumDecl             = 3
	ModelParserRULE_enumMemberDecl       = 4
	ModelParserRULE_classDecl            = 5
	ModelParserRULE_structDecl           = 6
	ModelParserRULE_structMemberDecl     = 7
	ModelParserRULE_attributeKind        = 8
	ModelParserRULE_typeReference        = 9
	ModelParserRULE_plainTypeReference   = 10
	ModelParserRULE_listTypeReference    = 11
	ModelParserRULE_mapTypeReference     = 12
	ModelParserRULE_resourceDecl         = 13
	ModelParserRULE_resourceMemberDecl   = 14
	ModelParserRULE_methodDecl           = 15
	ModelParserRULE_methodMemberDecl     = 16
	ModelParserRULE_methodParameterDecl  = 17
	ModelParserRULE_parameterDirection   = 18
	ModelParserRULE_locatorDecl          = 19
	ModelParserRULE_locatorMemberDecl    = 20
	ModelParserRULE_locatorTargetDecl    = 21
	ModelParserRULE_locatorVariableDecl  = 22
	ModelParserRULE_resourceReference    = 23
	ModelParserRULE_errorDecl            = 24
	ModelParserRULE_errorMemberDecl      = 25
	ModelParserRULE_errorCodeDecl        = 26
	ModelParserRULE_annotation           = 27
	ModelParserRULE_annotationParameters = 28
	ModelParserRULE_annotationParameter  = 29
	ModelParserRULE_literal              = 30
	ModelParserRULE_booleanLiteral       = 31
	ModelParserRULE_integerLiteral       = 32
	ModelParserRULE_stringLiteral        = 33
	ModelParserRULE_identifier           = 34
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *FileContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *ModelParser) File() (localctx IFileContext) {
	this := p
	_ = this

	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ModelParserRULE_file)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ModelParserCLASS)|(1<<ModelParserENUM)|(1<<ModelParserERROR)|(1<<ModelParserRESOURCE)|(1<<ModelParserSTRUCT)|(1<<ModelParserAT_SIGN))) != 0 {
		{
			p.SetState(70)
			p.Declaration()
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) TypeDecl() ITypeDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *DeclarationContext) ResourceDecl() IResourceDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResourceDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IResourceDeclContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *ModelParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ModelParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.TypeDecl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.ResourceDecl()
		}

	}

	return localctx
}

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetResult returns the result attribute.
	GetResult() *concepts.Type

	// SetResult sets the result attribute.
	SetResult(*concepts.Type)

	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	result *concepts.Type
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_typeDecl
	return p
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) GetResult() *concepts.Type { return s.result }

func (s *TypeDeclContext) SetResult(v *concepts.Type) { s.result = v }

func (s *TypeDeclContext) EnumDecl() IEnumDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumDeclContext)
}

func (s *TypeDeclContext) ClassDecl() IClassDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassDeclContext)
}

func (s *TypeDeclContext) StructDecl() IStructDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructDeclContext)
}

func (s *TypeDeclContext) ErrorDecl() IErrorDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IErrorDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IErrorDeclContext)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterTypeDecl(s)
	}
}

func (s *TypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitTypeDecl(s)
	}
}

func (p *ModelParser) TypeDecl() (localctx ITypeDeclContext) {
	this := p
	_ = this

	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ModelParserRULE_typeDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.EnumDecl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.ClassDecl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.StructDecl()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.ErrorDecl()
		}

	}

	return localctx
}

// IEnumDeclContext is an interface to support dynamic dispatch.
type IEnumDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_annotation returns the _annotation rule contexts.
	Get_annotation() IAnnotationContext

	// GetName returns the name rule contexts.
	GetName() IIdentifierContext

	// Get_enumMemberDecl returns the _enumMemberDecl rule contexts.
	Get_enumMemberDecl() IEnumMemberDeclContext

	// Set_annotation sets the _annotation rule contexts.
	Set_annotation(IAnnotationContext)

	// SetName sets the name rule contexts.
	SetName(IIdentifierContext)

	// Set_enumMemberDecl sets the _enumMemberDecl rule contexts.
	Set_enumMemberDecl(IEnumMemberDeclContext)

	// GetAnnotations returns the annotations rule context list.
	GetAnnotations() []IAnnotationContext

	// GetMembers returns the members rule context list.
	GetMembers() []IEnumMemberDeclContext

	// SetAnnotations sets the annotations rule context list.
	SetAnnotations([]IAnnotationContext)

	// SetMembers sets the members rule context list.
	SetMembers([]IEnumMemberDeclContext)

	// GetResult returns the result attribute.
	GetResult() *concepts.Type

	// SetResult sets the result attribute.
	SetResult(*concepts.Type)

	// IsEnumDeclContext differentiates from other interfaces.
	IsEnumDeclContext()
}

type EnumDeclContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	result          *concepts.Type
	_annotation     IAnnotationContext
	annotations     []IAnnotationContext
	name            IIdentifierContext
	_enumMemberDecl IEnumMemberDeclContext
	members         []IEnumMemberDeclContext
}

func NewEmptyEnumDeclContext() *EnumDeclContext {
	var p = new(EnumDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_enumDecl
	return p
}

func (*EnumDeclContext) IsEnumDeclContext() {}

func NewEnumDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDeclContext {
	var p = new(EnumDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_enumDecl

	return p
}

func (s *EnumDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDeclContext) Get_annotation() IAnnotationContext { return s._annotation }

func (s *EnumDeclContext) GetName() IIdentifierContext { return s.name }

func (s *EnumDeclContext) Get_enumMemberDecl() IEnumMemberDeclContext { return s._enumMemberDecl }

func (s *EnumDeclContext) Set_annotation(v IAnnotationContext) { s._annotation = v }

func (s *EnumDeclContext) SetName(v IIdentifierContext) { s.name = v }

func (s *EnumDeclContext) Set_enumMemberDecl(v IEnumMemberDeclContext) { s._enumMemberDecl = v }

func (s *EnumDeclContext) GetAnnotations() []IAnnotationContext { return s.annotations }

func (s *EnumDeclContext) GetMembers() []IEnumMemberDeclContext { return s.members }

func (s *EnumDeclContext) SetAnnotations(v []IAnnotationContext) { s.annotations = v }

func (s *EnumDeclContext) SetMembers(v []IEnumMemberDeclContext) { s.members = v }

func (s *EnumDeclContext) GetResult() *concepts.Type { return s.result }

func (s *EnumDeclContext) SetResult(v *concepts.Type) { s.result = v }

func (s *EnumDeclContext) ENUM() antlr.TerminalNode {
	return s.GetToken(ModelParserENUM, 0)
}

func (s *EnumDeclContext) LEFT_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserLEFT_CURLY_BRACKET, 0)
}

func (s *EnumDeclContext) RIGHT_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserRIGHT_CURLY_BRACKET, 0)
}

func (s *EnumDeclContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *EnumDeclContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *EnumDeclContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *EnumDeclContext) AllEnumMemberDecl() []IEnumMemberDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumMemberDeclContext)(nil)).Elem())
	var tst = make([]IEnumMemberDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumMemberDeclContext)
		}
	}

	return tst
}

func (s *EnumDeclContext) EnumMemberDecl(i int) IEnumMemberDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumMemberDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumMemberDeclContext)
}

func (s *EnumDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterEnumDecl(s)
	}
}

func (s *EnumDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitEnumDecl(s)
	}
}

func (p *ModelParser) EnumDecl() (localctx IEnumDeclContext) {
	this := p
	_ = this

	localctx = NewEnumDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ModelParserRULE_enumDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ModelParserAT_SIGN {
		{
			p.SetState(86)

			var _x = p.Annotation()

			localctx.(*EnumDeclContext)._annotation = _x
		}
		localctx.(*EnumDeclContext).annotations = append(localctx.(*EnumDeclContext).annotations, localctx.(*EnumDeclContext)._annotation)

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(92)
		p.Match(ModelParserENUM)
	}
	{
		p.SetState(93)

		var _x = p.Identifier()

		localctx.(*EnumDeclContext).name = _x
	}
	{
		p.SetState(94)
		p.Match(ModelParserLEFT_CURLY_BRACKET)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ModelParserVALUE)|(1<<ModelParserAT_SIGN)|(1<<ModelParserIDENTIFIER))) != 0 {
		{
			p.SetState(95)

			var _x = p.EnumMemberDecl()

			localctx.(*EnumDeclContext)._enumMemberDecl = _x
		}
		localctx.(*EnumDeclContext).members = append(localctx.(*EnumDeclContext).members, localctx.(*EnumDeclContext)._enumMemberDecl)

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(101)
		p.Match(ModelParserRIGHT_CURLY_BRACKET)
	}

	return localctx
}

// IEnumMemberDeclContext is an interface to support dynamic dispatch.
type IEnumMemberDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_annotation returns the _annotation rule contexts.
	Get_annotation() IAnnotationContext

	// GetName returns the name rule contexts.
	GetName() IIdentifierContext

	// Set_annotation sets the _annotation rule contexts.
	Set_annotation(IAnnotationContext)

	// SetName sets the name rule contexts.
	SetName(IIdentifierContext)

	// GetAnnotations returns the annotations rule context list.
	GetAnnotations() []IAnnotationContext

	// SetAnnotations sets the annotations rule context list.
	SetAnnotations([]IAnnotationContext)

	// GetResult returns the result attribute.
	GetResult() *concepts.EnumValue

	// SetResult sets the result attribute.
	SetResult(*concepts.EnumValue)

	// IsEnumMemberDeclContext differentiates from other interfaces.
	IsEnumMemberDeclContext()
}

type EnumMemberDeclContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	result      *concepts.EnumValue
	_annotation IAnnotationContext
	annotations []IAnnotationContext
	name        IIdentifierContext
}

func NewEmptyEnumMemberDeclContext() *EnumMemberDeclContext {
	var p = new(EnumMemberDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_enumMemberDecl
	return p
}

func (*EnumMemberDeclContext) IsEnumMemberDeclContext() {}

func NewEnumMemberDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumMemberDeclContext {
	var p = new(EnumMemberDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_enumMemberDecl

	return p
}

func (s *EnumMemberDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumMemberDeclContext) Get_annotation() IAnnotationContext { return s._annotation }

func (s *EnumMemberDeclContext) GetName() IIdentifierContext { return s.name }

func (s *EnumMemberDeclContext) Set_annotation(v IAnnotationContext) { s._annotation = v }

func (s *EnumMemberDeclContext) SetName(v IIdentifierContext) { s.name = v }

func (s *EnumMemberDeclContext) GetAnnotations() []IAnnotationContext { return s.annotations }

func (s *EnumMemberDeclContext) SetAnnotations(v []IAnnotationContext) { s.annotations = v }

func (s *EnumMemberDeclContext) GetResult() *concepts.EnumValue { return s.result }

func (s *EnumMemberDeclContext) SetResult(v *concepts.EnumValue) { s.result = v }

func (s *EnumMemberDeclContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *EnumMemberDeclContext) VALUE() antlr.TerminalNode {
	return s.GetToken(ModelParserVALUE, 0)
}

func (s *EnumMemberDeclContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *EnumMemberDeclContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *EnumMemberDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumMemberDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumMemberDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterEnumMemberDecl(s)
	}
}

func (s *EnumMemberDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitEnumMemberDecl(s)
	}
}

func (p *ModelParser) EnumMemberDecl() (localctx IEnumMemberDeclContext) {
	this := p
	_ = this

	localctx = NewEnumMemberDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ModelParserRULE_enumMemberDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ModelParserAT_SIGN {
		{
			p.SetState(103)

			var _x = p.Annotation()

			localctx.(*EnumMemberDeclContext)._annotation = _x
		}
		localctx.(*EnumMemberDeclContext).annotations = append(localctx.(*EnumMemberDeclContext).annotations, localctx.(*EnumMemberDeclContext)._annotation)

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ModelParserVALUE {
		{
			p.SetState(109)
			p.Match(ModelParserVALUE)
		}

	}
	{
		p.SetState(112)

		var _x = p.Identifier()

		localctx.(*EnumMemberDeclContext).name = _x
	}

	return localctx
}

// IClassDeclContext is an interface to support dynamic dispatch.
type IClassDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_annotation returns the _annotation rule contexts.
	Get_annotation() IAnnotationContext

	// GetName returns the name rule contexts.
	GetName() IIdentifierContext

	// Get_structMemberDecl returns the _structMemberDecl rule contexts.
	Get_structMemberDecl() IStructMemberDeclContext

	// Set_annotation sets the _annotation rule contexts.
	Set_annotation(IAnnotationContext)

	// SetName sets the name rule contexts.
	SetName(IIdentifierContext)

	// Set_structMemberDecl sets the _structMemberDecl rule contexts.
	Set_structMemberDecl(IStructMemberDeclContext)

	// GetAnnotations returns the annotations rule context list.
	GetAnnotations() []IAnnotationContext

	// GetMembers returns the members rule context list.
	GetMembers() []IStructMemberDeclContext

	// SetAnnotations sets the annotations rule context list.
	SetAnnotations([]IAnnotationContext)

	// SetMembers sets the members rule context list.
	SetMembers([]IStructMemberDeclContext)

	// GetResult returns the result attribute.
	GetResult() *concepts.Type

	// SetResult sets the result attribute.
	SetResult(*concepts.Type)

	// IsClassDeclContext differentiates from other interfaces.
	IsClassDeclContext()
}

type ClassDeclContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	result            *concepts.Type
	_annotation       IAnnotationContext
	annotations       []IAnnotationContext
	name              IIdentifierContext
	_structMemberDecl IStructMemberDeclContext
	members           []IStructMemberDeclContext
}

func NewEmptyClassDeclContext() *ClassDeclContext {
	var p = new(ClassDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_classDecl
	return p
}

func (*ClassDeclContext) IsClassDeclContext() {}

func NewClassDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDeclContext {
	var p = new(ClassDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_classDecl

	return p
}

func (s *ClassDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDeclContext) Get_annotation() IAnnotationContext { return s._annotation }

func (s *ClassDeclContext) GetName() IIdentifierContext { return s.name }

func (s *ClassDeclContext) Get_structMemberDecl() IStructMemberDeclContext {
	return s._structMemberDecl
}

func (s *ClassDeclContext) Set_annotation(v IAnnotationContext) { s._annotation = v }

func (s *ClassDeclContext) SetName(v IIdentifierContext) { s.name = v }

func (s *ClassDeclContext) Set_structMemberDecl(v IStructMemberDeclContext) { s._structMemberDecl = v }

func (s *ClassDeclContext) GetAnnotations() []IAnnotationContext { return s.annotations }

func (s *ClassDeclContext) GetMembers() []IStructMemberDeclContext { return s.members }

func (s *ClassDeclContext) SetAnnotations(v []IAnnotationContext) { s.annotations = v }

func (s *ClassDeclContext) SetMembers(v []IStructMemberDeclContext) { s.members = v }

func (s *ClassDeclContext) GetResult() *concepts.Type { return s.result }

func (s *ClassDeclContext) SetResult(v *concepts.Type) { s.result = v }

func (s *ClassDeclContext) CLASS() antlr.TerminalNode {
	return s.GetToken(ModelParserCLASS, 0)
}

func (s *ClassDeclContext) LEFT_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserLEFT_CURLY_BRACKET, 0)
}

func (s *ClassDeclContext) RIGHT_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserRIGHT_CURLY_BRACKET, 0)
}

func (s *ClassDeclContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ClassDeclContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *ClassDeclContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *ClassDeclContext) AllStructMemberDecl() []IStructMemberDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStructMemberDeclContext)(nil)).Elem())
	var tst = make([]IStructMemberDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStructMemberDeclContext)
		}
	}

	return tst
}

func (s *ClassDeclContext) StructMemberDecl(i int) IStructMemberDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructMemberDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStructMemberDeclContext)
}

func (s *ClassDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterClassDecl(s)
	}
}

func (s *ClassDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitClassDecl(s)
	}
}

func (p *ModelParser) ClassDecl() (localctx IClassDeclContext) {
	this := p
	_ = this

	localctx = NewClassDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ModelParserRULE_classDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ModelParserAT_SIGN {
		{
			p.SetState(114)

			var _x = p.Annotation()

			localctx.(*ClassDeclContext)._annotation = _x
		}
		localctx.(*ClassDeclContext).annotations = append(localctx.(*ClassDeclContext).annotations, localctx.(*ClassDeclContext)._annotation)

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(120)
		p.Match(ModelParserCLASS)
	}
	{
		p.SetState(121)

		var _x = p.Identifier()

		localctx.(*ClassDeclContext).name = _x
	}
	{
		p.SetState(122)
		p.Match(ModelParserLEFT_CURLY_BRACKET)
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ModelParserATTRIBUTE)|(1<<ModelParserLINK)|(1<<ModelParserAT_SIGN)|(1<<ModelParserIDENTIFIER))) != 0 {
		{
			p.SetState(123)

			var _x = p.StructMemberDecl()

			localctx.(*ClassDeclContext)._structMemberDecl = _x
		}
		localctx.(*ClassDeclContext).members = append(localctx.(*ClassDeclContext).members, localctx.(*ClassDeclContext)._structMemberDecl)

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(129)
		p.Match(ModelParserRIGHT_CURLY_BRACKET)
	}

	return localctx
}

// IStructDeclContext is an interface to support dynamic dispatch.
type IStructDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_annotation returns the _annotation rule contexts.
	Get_annotation() IAnnotationContext

	// GetName returns the name rule contexts.
	GetName() IIdentifierContext

	// Get_structMemberDecl returns the _structMemberDecl rule contexts.
	Get_structMemberDecl() IStructMemberDeclContext

	// Set_annotation sets the _annotation rule contexts.
	Set_annotation(IAnnotationContext)

	// SetName sets the name rule contexts.
	SetName(IIdentifierContext)

	// Set_structMemberDecl sets the _structMemberDecl rule contexts.
	Set_structMemberDecl(IStructMemberDeclContext)

	// GetAnnotations returns the annotations rule context list.
	GetAnnotations() []IAnnotationContext

	// GetMembers returns the members rule context list.
	GetMembers() []IStructMemberDeclContext

	// SetAnnotations sets the annotations rule context list.
	SetAnnotations([]IAnnotationContext)

	// SetMembers sets the members rule context list.
	SetMembers([]IStructMemberDeclContext)

	// GetResult returns the result attribute.
	GetResult() *concepts.Type

	// SetResult sets the result attribute.
	SetResult(*concepts.Type)

	// IsStructDeclContext differentiates from other interfaces.
	IsStructDeclContext()
}

type StructDeclContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	result            *concepts.Type
	_annotation       IAnnotationContext
	annotations       []IAnnotationContext
	name              IIdentifierContext
	_structMemberDecl IStructMemberDeclContext
	members           []IStructMemberDeclContext
}

func NewEmptyStructDeclContext() *StructDeclContext {
	var p = new(StructDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_structDecl
	return p
}

func (*StructDeclContext) IsStructDeclContext() {}

func NewStructDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclContext {
	var p = new(StructDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_structDecl

	return p
}

func (s *StructDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclContext) Get_annotation() IAnnotationContext { return s._annotation }

func (s *StructDeclContext) GetName() IIdentifierContext { return s.name }

func (s *StructDeclContext) Get_structMemberDecl() IStructMemberDeclContext {
	return s._structMemberDecl
}

func (s *StructDeclContext) Set_annotation(v IAnnotationContext) { s._annotation = v }

func (s *StructDeclContext) SetName(v IIdentifierContext) { s.name = v }

func (s *StructDeclContext) Set_structMemberDecl(v IStructMemberDeclContext) { s._structMemberDecl = v }

func (s *StructDeclContext) GetAnnotations() []IAnnotationContext { return s.annotations }

func (s *StructDeclContext) GetMembers() []IStructMemberDeclContext { return s.members }

func (s *StructDeclContext) SetAnnotations(v []IAnnotationContext) { s.annotations = v }

func (s *StructDeclContext) SetMembers(v []IStructMemberDeclContext) { s.members = v }

func (s *StructDeclContext) GetResult() *concepts.Type { return s.result }

func (s *StructDeclContext) SetResult(v *concepts.Type) { s.result = v }

func (s *StructDeclContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(ModelParserSTRUCT, 0)
}

func (s *StructDeclContext) LEFT_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserLEFT_CURLY_BRACKET, 0)
}

func (s *StructDeclContext) RIGHT_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserRIGHT_CURLY_BRACKET, 0)
}

func (s *StructDeclContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *StructDeclContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *StructDeclContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *StructDeclContext) AllStructMemberDecl() []IStructMemberDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStructMemberDeclContext)(nil)).Elem())
	var tst = make([]IStructMemberDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStructMemberDeclContext)
		}
	}

	return tst
}

func (s *StructDeclContext) StructMemberDecl(i int) IStructMemberDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructMemberDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStructMemberDeclContext)
}

func (s *StructDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterStructDecl(s)
	}
}

func (s *StructDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitStructDecl(s)
	}
}

func (p *ModelParser) StructDecl() (localctx IStructDeclContext) {
	this := p
	_ = this

	localctx = NewStructDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ModelParserRULE_structDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ModelParserAT_SIGN {
		{
			p.SetState(131)

			var _x = p.Annotation()

			localctx.(*StructDeclContext)._annotation = _x
		}
		localctx.(*StructDeclContext).annotations = append(localctx.(*StructDeclContext).annotations, localctx.(*StructDeclContext)._annotation)

		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(137)
		p.Match(ModelParserSTRUCT)
	}
	{
		p.SetState(138)

		var _x = p.Identifier()

		localctx.(*StructDeclContext).name = _x
	}
	{
		p.SetState(139)
		p.Match(ModelParserLEFT_CURLY_BRACKET)
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ModelParserATTRIBUTE)|(1<<ModelParserLINK)|(1<<ModelParserAT_SIGN)|(1<<ModelParserIDENTIFIER))) != 0 {
		{
			p.SetState(140)

			var _x = p.StructMemberDecl()

			localctx.(*StructDeclContext)._structMemberDecl = _x
		}
		localctx.(*StructDeclContext).members = append(localctx.(*StructDeclContext).members, localctx.(*StructDeclContext)._structMemberDecl)

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(146)
		p.Match(ModelParserRIGHT_CURLY_BRACKET)
	}

	return localctx
}

// IStructMemberDeclContext is an interface to support dynamic dispatch.
type IStructMemberDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_annotation returns the _annotation rule contexts.
	Get_annotation() IAnnotationContext

	// GetKind returns the kind rule contexts.
	GetKind() IAttributeKindContext

	// GetName returns the name rule contexts.
	GetName() IIdentifierContext

	// GetReference returns the reference rule contexts.
	GetReference() ITypeReferenceContext

	// Set_annotation sets the _annotation rule contexts.
	Set_annotation(IAnnotationContext)

	// SetKind sets the kind rule contexts.
	SetKind(IAttributeKindContext)

	// SetName sets the name rule contexts.
	SetName(IIdentifierContext)

	// SetReference sets the reference rule contexts.
	SetReference(ITypeReferenceContext)

	// GetAnnotations returns the annotations rule context list.
	GetAnnotations() []IAnnotationContext

	// SetAnnotations sets the annotations rule context list.
	SetAnnotations([]IAnnotationContext)

	// GetResult returns the result attribute.
	GetResult() *concepts.Attribute

	// SetResult sets the result attribute.
	SetResult(*concepts.Attribute)

	// IsStructMemberDeclContext differentiates from other interfaces.
	IsStructMemberDeclContext()
}

type StructMemberDeclContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	result      *concepts.Attribute
	_annotation IAnnotationContext
	annotations []IAnnotationContext
	kind        IAttributeKindContext
	name        IIdentifierContext
	reference   ITypeReferenceContext
}

func NewEmptyStructMemberDeclContext() *StructMemberDeclContext {
	var p = new(StructMemberDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_structMemberDecl
	return p
}

func (*StructMemberDeclContext) IsStructMemberDeclContext() {}

func NewStructMemberDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructMemberDeclContext {
	var p = new(StructMemberDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_structMemberDecl

	return p
}

func (s *StructMemberDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StructMemberDeclContext) Get_annotation() IAnnotationContext { return s._annotation }

func (s *StructMemberDeclContext) GetKind() IAttributeKindContext { return s.kind }

func (s *StructMemberDeclContext) GetName() IIdentifierContext { return s.name }

func (s *StructMemberDeclContext) GetReference() ITypeReferenceContext { return s.reference }

func (s *StructMemberDeclContext) Set_annotation(v IAnnotationContext) { s._annotation = v }

func (s *StructMemberDeclContext) SetKind(v IAttributeKindContext) { s.kind = v }

func (s *StructMemberDeclContext) SetName(v IIdentifierContext) { s.name = v }

func (s *StructMemberDeclContext) SetReference(v ITypeReferenceContext) { s.reference = v }

func (s *StructMemberDeclContext) GetAnnotations() []IAnnotationContext { return s.annotations }

func (s *StructMemberDeclContext) SetAnnotations(v []IAnnotationContext) { s.annotations = v }

func (s *StructMemberDeclContext) GetResult() *concepts.Attribute { return s.result }

func (s *StructMemberDeclContext) SetResult(v *concepts.Attribute) { s.result = v }

func (s *StructMemberDeclContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *StructMemberDeclContext) TypeReference() ITypeReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *StructMemberDeclContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *StructMemberDeclContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *StructMemberDeclContext) AttributeKind() IAttributeKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributeKindContext)
}

func (s *StructMemberDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructMemberDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructMemberDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterStructMemberDecl(s)
	}
}

func (s *StructMemberDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitStructMemberDecl(s)
	}
}

func (p *ModelParser) StructMemberDecl() (localctx IStructMemberDeclContext) {
	this := p
	_ = this

	localctx = NewStructMemberDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ModelParserRULE_structMemberDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ModelParserAT_SIGN {
		{
			p.SetState(148)

			var _x = p.Annotation()

			localctx.(*StructMemberDeclContext)._annotation = _x
		}
		localctx.(*StructMemberDeclContext).annotations = append(localctx.(*StructMemberDeclContext).annotations, localctx.(*StructMemberDeclContext)._annotation)

		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ModelParserATTRIBUTE || _la == ModelParserLINK {
		{
			p.SetState(154)

			var _x = p.AttributeKind()

			localctx.(*StructMemberDeclContext).kind = _x
		}

	}
	{
		p.SetState(157)

		var _x = p.Identifier()

		localctx.(*StructMemberDeclContext).name = _x
	}
	{
		p.SetState(158)

		var _x = p.TypeReference()

		localctx.(*StructMemberDeclContext).reference = _x
	}

	return localctx
}

// IAttributeKindContext is an interface to support dynamic dispatch.
type IAttributeKindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetResult returns the result attribute.
	GetResult() int

	// SetResult sets the result attribute.
	SetResult(int)

	// IsAttributeKindContext differentiates from other interfaces.
	IsAttributeKindContext()
}

type AttributeKindContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	result int
}

func NewEmptyAttributeKindContext() *AttributeKindContext {
	var p = new(AttributeKindContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_attributeKind
	return p
}

func (*AttributeKindContext) IsAttributeKindContext() {}

func NewAttributeKindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeKindContext {
	var p = new(AttributeKindContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_attributeKind

	return p
}

func (s *AttributeKindContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeKindContext) GetResult() int { return s.result }

func (s *AttributeKindContext) SetResult(v int) { s.result = v }

func (s *AttributeKindContext) ATTRIBUTE() antlr.TerminalNode {
	return s.GetToken(ModelParserATTRIBUTE, 0)
}

func (s *AttributeKindContext) LINK() antlr.TerminalNode {
	return s.GetToken(ModelParserLINK, 0)
}

func (s *AttributeKindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeKindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeKindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterAttributeKind(s)
	}
}

func (s *AttributeKindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitAttributeKind(s)
	}
}

func (p *ModelParser) AttributeKind() (localctx IAttributeKindContext) {
	this := p
	_ = this

	localctx = NewAttributeKindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ModelParserRULE_attributeKind)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ModelParserATTRIBUTE || _la == ModelParserLINK) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITypeReferenceContext is an interface to support dynamic dispatch.
type ITypeReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPlain returns the plain rule contexts.
	GetPlain() IPlainTypeReferenceContext

	// GetList returns the list rule contexts.
	GetList() IListTypeReferenceContext

	// GetMp returns the mp rule contexts.
	GetMp() IMapTypeReferenceContext

	// SetPlain sets the plain rule contexts.
	SetPlain(IPlainTypeReferenceContext)

	// SetList sets the list rule contexts.
	SetList(IListTypeReferenceContext)

	// SetMp sets the mp rule contexts.
	SetMp(IMapTypeReferenceContext)

	// GetResult returns the result attribute.
	GetResult() *concepts.Type

	// SetResult sets the result attribute.
	SetResult(*concepts.Type)

	// IsTypeReferenceContext differentiates from other interfaces.
	IsTypeReferenceContext()
}

type TypeReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	result *concepts.Type
	plain  IPlainTypeReferenceContext
	list   IListTypeReferenceContext
	mp     IMapTypeReferenceContext
}

func NewEmptyTypeReferenceContext() *TypeReferenceContext {
	var p = new(TypeReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_typeReference
	return p
}

func (*TypeReferenceContext) IsTypeReferenceContext() {}

func NewTypeReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeReferenceContext {
	var p = new(TypeReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_typeReference

	return p
}

func (s *TypeReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeReferenceContext) GetPlain() IPlainTypeReferenceContext { return s.plain }

func (s *TypeReferenceContext) GetList() IListTypeReferenceContext { return s.list }

func (s *TypeReferenceContext) GetMp() IMapTypeReferenceContext { return s.mp }

func (s *TypeReferenceContext) SetPlain(v IPlainTypeReferenceContext) { s.plain = v }

func (s *TypeReferenceContext) SetList(v IListTypeReferenceContext) { s.list = v }

func (s *TypeReferenceContext) SetMp(v IMapTypeReferenceContext) { s.mp = v }

func (s *TypeReferenceContext) GetResult() *concepts.Type { return s.result }

func (s *TypeReferenceContext) SetResult(v *concepts.Type) { s.result = v }

func (s *TypeReferenceContext) PlainTypeReference() IPlainTypeReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlainTypeReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPlainTypeReferenceContext)
}

func (s *TypeReferenceContext) ListTypeReference() IListTypeReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListTypeReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListTypeReferenceContext)
}

func (s *TypeReferenceContext) MapTypeReference() IMapTypeReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapTypeReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapTypeReferenceContext)
}

func (s *TypeReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterTypeReference(s)
	}
}

func (s *TypeReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitTypeReference(s)
	}
}

func (p *ModelParser) TypeReference() (localctx ITypeReferenceContext) {
	this := p
	_ = this

	localctx = NewTypeReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ModelParserRULE_typeReference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)

			var _x = p.PlainTypeReference()

			localctx.(*TypeReferenceContext).plain = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)

			var _x = p.ListTypeReference()

			localctx.(*TypeReferenceContext).list = _x
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(164)

			var _x = p.MapTypeReference()

			localctx.(*TypeReferenceContext).mp = _x
		}

	}

	return localctx
}

// IPlainTypeReferenceContext is an interface to support dynamic dispatch.
type IPlainTypeReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IIdentifierContext

	// SetName sets the name rule contexts.
	SetName(IIdentifierContext)

	// GetResult returns the result attribute.
	GetResult() *concepts.Type

	// SetResult sets the result attribute.
	SetResult(*concepts.Type)

	// IsPlainTypeReferenceContext differentiates from other interfaces.
	IsPlainTypeReferenceContext()
}

type PlainTypeReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	result *concepts.Type
	name   IIdentifierContext
}

func NewEmptyPlainTypeReferenceContext() *PlainTypeReferenceContext {
	var p = new(PlainTypeReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_plainTypeReference
	return p
}

func (*PlainTypeReferenceContext) IsPlainTypeReferenceContext() {}

func NewPlainTypeReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlainTypeReferenceContext {
	var p = new(PlainTypeReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_plainTypeReference

	return p
}

func (s *PlainTypeReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *PlainTypeReferenceContext) GetName() IIdentifierContext { return s.name }

func (s *PlainTypeReferenceContext) SetName(v IIdentifierContext) { s.name = v }

func (s *PlainTypeReferenceContext) GetResult() *concepts.Type { return s.result }

func (s *PlainTypeReferenceContext) SetResult(v *concepts.Type) { s.result = v }

func (s *PlainTypeReferenceContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *PlainTypeReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlainTypeReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlainTypeReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterPlainTypeReference(s)
	}
}

func (s *PlainTypeReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitPlainTypeReference(s)
	}
}

func (p *ModelParser) PlainTypeReference() (localctx IPlainTypeReferenceContext) {
	this := p
	_ = this

	localctx = NewPlainTypeReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ModelParserRULE_plainTypeReference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)

		var _x = p.Identifier()

		localctx.(*PlainTypeReferenceContext).name = _x
	}

	return localctx
}

// IListTypeReferenceContext is an interface to support dynamic dispatch.
type IListTypeReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetElement returns the element rule contexts.
	GetElement() IIdentifierContext

	// SetElement sets the element rule contexts.
	SetElement(IIdentifierContext)

	// GetResult returns the result attribute.
	GetResult() *concepts.Type

	// SetResult sets the result attribute.
	SetResult(*concepts.Type)

	// IsListTypeReferenceContext differentiates from other interfaces.
	IsListTypeReferenceContext()
}

type ListTypeReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	result  *concepts.Type
	element IIdentifierContext
}

func NewEmptyListTypeReferenceContext() *ListTypeReferenceContext {
	var p = new(ListTypeReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_listTypeReference
	return p
}

func (*ListTypeReferenceContext) IsListTypeReferenceContext() {}

func NewListTypeReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListTypeReferenceContext {
	var p = new(ListTypeReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_listTypeReference

	return p
}

func (s *ListTypeReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ListTypeReferenceContext) GetElement() IIdentifierContext { return s.element }

func (s *ListTypeReferenceContext) SetElement(v IIdentifierContext) { s.element = v }

func (s *ListTypeReferenceContext) GetResult() *concepts.Type { return s.result }

func (s *ListTypeReferenceContext) SetResult(v *concepts.Type) { s.result = v }

func (s *ListTypeReferenceContext) LEFT_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserLEFT_SQUARE_BRACKET, 0)
}

func (s *ListTypeReferenceContext) RIGHT_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserRIGHT_SQUARE_BRACKET, 0)
}

func (s *ListTypeReferenceContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ListTypeReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListTypeReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListTypeReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterListTypeReference(s)
	}
}

func (s *ListTypeReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitListTypeReference(s)
	}
}

func (p *ModelParser) ListTypeReference() (localctx IListTypeReferenceContext) {
	this := p
	_ = this

	localctx = NewListTypeReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ModelParserRULE_listTypeReference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(ModelParserLEFT_SQUARE_BRACKET)
	}
	{
		p.SetState(170)
		p.Match(ModelParserRIGHT_SQUARE_BRACKET)
	}
	{
		p.SetState(171)

		var _x = p.Identifier()

		localctx.(*ListTypeReferenceContext).element = _x
	}

	return localctx
}

// IMapTypeReferenceContext is an interface to support dynamic dispatch.
type IMapTypeReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIndex returns the index rule contexts.
	GetIndex() IIdentifierContext

	// GetElement returns the element rule contexts.
	GetElement() IIdentifierContext

	// SetIndex sets the index rule contexts.
	SetIndex(IIdentifierContext)

	// SetElement sets the element rule contexts.
	SetElement(IIdentifierContext)

	// GetResult returns the result attribute.
	GetResult() *concepts.Type

	// SetResult sets the result attribute.
	SetResult(*concepts.Type)

	// IsMapTypeReferenceContext differentiates from other interfaces.
	IsMapTypeReferenceContext()
}

type MapTypeReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	result  *concepts.Type
	index   IIdentifierContext
	element IIdentifierContext
}

func NewEmptyMapTypeReferenceContext() *MapTypeReferenceContext {
	var p = new(MapTypeReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_mapTypeReference
	return p
}

func (*MapTypeReferenceContext) IsMapTypeReferenceContext() {}

func NewMapTypeReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeReferenceContext {
	var p = new(MapTypeReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_mapTypeReference

	return p
}

func (s *MapTypeReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *MapTypeReferenceContext) GetIndex() IIdentifierContext { return s.index }

func (s *MapTypeReferenceContext) GetElement() IIdentifierContext { return s.element }

func (s *MapTypeReferenceContext) SetIndex(v IIdentifierContext) { s.index = v }

func (s *MapTypeReferenceContext) SetElement(v IIdentifierContext) { s.element = v }

func (s *MapTypeReferenceContext) GetResult() *concepts.Type { return s.result }

func (s *MapTypeReferenceContext) SetResult(v *concepts.Type) { s.result = v }

func (s *MapTypeReferenceContext) LEFT_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserLEFT_SQUARE_BRACKET, 0)
}

func (s *MapTypeReferenceContext) RIGHT_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserRIGHT_SQUARE_BRACKET, 0)
}

func (s *MapTypeReferenceContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *MapTypeReferenceContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MapTypeReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTypeReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapTypeReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterMapTypeReference(s)
	}
}

func (s *MapTypeReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitMapTypeReference(s)
	}
}

func (p *ModelParser) MapTypeReference() (localctx IMapTypeReferenceContext) {
	this := p
	_ = this

	localctx = NewMapTypeReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ModelParserRULE_mapTypeReference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(ModelParserLEFT_SQUARE_BRACKET)
	}
	{
		p.SetState(174)

		var _x = p.Identifier()

		localctx.(*MapTypeReferenceContext).index = _x
	}
	{
		p.SetState(175)
		p.Match(ModelParserRIGHT_SQUARE_BRACKET)
	}
	{
		p.SetState(176)

		var _x = p.Identifier()

		localctx.(*MapTypeReferenceContext).element = _x
	}

	return localctx
}

// IResourceDeclContext is an interface to support dynamic dispatch.
type IResourceDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_annotation returns the _annotation rule contexts.
	Get_annotation() IAnnotationContext

	// GetName returns the name rule contexts.
	GetName() IIdentifierContext

	// Get_resourceMemberDecl returns the _resourceMemberDecl rule contexts.
	Get_resourceMemberDecl() IResourceMemberDeclContext

	// Set_annotation sets the _annotation rule contexts.
	Set_annotation(IAnnotationContext)

	// SetName sets the name rule contexts.
	SetName(IIdentifierContext)

	// Set_resourceMemberDecl sets the _resourceMemberDecl rule contexts.
	Set_resourceMemberDecl(IResourceMemberDeclContext)

	// GetAnnotations returns the annotations rule context list.
	GetAnnotations() []IAnnotationContext

	// GetMembers returns the members rule context list.
	GetMembers() []IResourceMemberDeclContext

	// SetAnnotations sets the annotations rule context list.
	SetAnnotations([]IAnnotationContext)

	// SetMembers sets the members rule context list.
	SetMembers([]IResourceMemberDeclContext)

	// GetResult returns the result attribute.
	GetResult() *concepts.Resource

	// SetResult sets the result attribute.
	SetResult(*concepts.Resource)

	// IsResourceDeclContext differentiates from other interfaces.
	IsResourceDeclContext()
}

type ResourceDeclContext struct {
	*antlr.BaseParserRuleContext
	parser              antlr.Parser
	result              *concepts.Resource
	_annotation         IAnnotationContext
	annotations         []IAnnotationContext
	name                IIdentifierContext
	_resourceMemberDecl IResourceMemberDeclContext
	members             []IResourceMemberDeclContext
}

func NewEmptyResourceDeclContext() *ResourceDeclContext {
	var p = new(ResourceDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_resourceDecl
	return p
}

func (*ResourceDeclContext) IsResourceDeclContext() {}

func NewResourceDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourceDeclContext {
	var p = new(ResourceDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_resourceDecl

	return p
}

func (s *ResourceDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourceDeclContext) Get_annotation() IAnnotationContext { return s._annotation }

func (s *ResourceDeclContext) GetName() IIdentifierContext { return s.name }

func (s *ResourceDeclContext) Get_resourceMemberDecl() IResourceMemberDeclContext {
	return s._resourceMemberDecl
}

func (s *ResourceDeclContext) Set_annotation(v IAnnotationContext) { s._annotation = v }

func (s *ResourceDeclContext) SetName(v IIdentifierContext) { s.name = v }

func (s *ResourceDeclContext) Set_resourceMemberDecl(v IResourceMemberDeclContext) {
	s._resourceMemberDecl = v
}

func (s *ResourceDeclContext) GetAnnotations() []IAnnotationContext { return s.annotations }

func (s *ResourceDeclContext) GetMembers() []IResourceMemberDeclContext { return s.members }

func (s *ResourceDeclContext) SetAnnotations(v []IAnnotationContext) { s.annotations = v }

func (s *ResourceDeclContext) SetMembers(v []IResourceMemberDeclContext) { s.members = v }

func (s *ResourceDeclContext) GetResult() *concepts.Resource { return s.result }

func (s *ResourceDeclContext) SetResult(v *concepts.Resource) { s.result = v }

func (s *ResourceDeclContext) RESOURCE() antlr.TerminalNode {
	return s.GetToken(ModelParserRESOURCE, 0)
}

func (s *ResourceDeclContext) LEFT_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserLEFT_CURLY_BRACKET, 0)
}

func (s *ResourceDeclContext) RIGHT_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserRIGHT_CURLY_BRACKET, 0)
}

func (s *ResourceDeclContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ResourceDeclContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *ResourceDeclContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *ResourceDeclContext) AllResourceMemberDecl() []IResourceMemberDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IResourceMemberDeclContext)(nil)).Elem())
	var tst = make([]IResourceMemberDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IResourceMemberDeclContext)
		}
	}

	return tst
}

func (s *ResourceDeclContext) ResourceMemberDecl(i int) IResourceMemberDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResourceMemberDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IResourceMemberDeclContext)
}

func (s *ResourceDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourceDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourceDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterResourceDecl(s)
	}
}

func (s *ResourceDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitResourceDecl(s)
	}
}

func (p *ModelParser) ResourceDecl() (localctx IResourceDeclContext) {
	this := p
	_ = this

	localctx = NewResourceDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ModelParserRULE_resourceDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ModelParserAT_SIGN {
		{
			p.SetState(178)

			var _x = p.Annotation()

			localctx.(*ResourceDeclContext)._annotation = _x
		}
		localctx.(*ResourceDeclContext).annotations = append(localctx.(*ResourceDeclContext).annotations, localctx.(*ResourceDeclContext)._annotation)

		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(184)
		p.Match(ModelParserRESOURCE)
	}
	{
		p.SetState(185)

		var _x = p.Identifier()

		localctx.(*ResourceDeclContext).name = _x
	}
	{
		p.SetState(186)
		p.Match(ModelParserLEFT_CURLY_BRACKET)
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ModelParserLOCATOR)|(1<<ModelParserMETHOD)|(1<<ModelParserAT_SIGN)|(1<<ModelParserIDENTIFIER))) != 0 {
		{
			p.SetState(187)

			var _x = p.ResourceMemberDecl()

			localctx.(*ResourceDeclContext)._resourceMemberDecl = _x
		}
		localctx.(*ResourceDeclContext).members = append(localctx.(*ResourceDeclContext).members, localctx.(*ResourceDeclContext)._resourceMemberDecl)

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(193)
		p.Match(ModelParserRIGHT_CURLY_BRACKET)
	}

	return localctx
}

// IResourceMemberDeclContext is an interface to support dynamic dispatch.
type IResourceMemberDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetResult returns the result attribute.
	GetResult() interface{}

	// SetResult sets the result attribute.
	SetResult(interface{})

	// IsResourceMemberDeclContext differentiates from other interfaces.
	IsResourceMemberDeclContext()
}

type ResourceMemberDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	result interface{}
}

func NewEmptyResourceMemberDeclContext() *ResourceMemberDeclContext {
	var p = new(ResourceMemberDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_resourceMemberDecl
	return p
}

func (*ResourceMemberDeclContext) IsResourceMemberDeclContext() {}

func NewResourceMemberDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourceMemberDeclContext {
	var p = new(ResourceMemberDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_resourceMemberDecl

	return p
}

func (s *ResourceMemberDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourceMemberDeclContext) GetResult() interface{} { return s.result }

func (s *ResourceMemberDeclContext) SetResult(v interface{}) { s.result = v }

func (s *ResourceMemberDeclContext) MethodDecl() IMethodDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodDeclContext)
}

func (s *ResourceMemberDeclContext) LocatorDecl() ILocatorDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocatorDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocatorDeclContext)
}

func (s *ResourceMemberDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourceMemberDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourceMemberDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterResourceMemberDecl(s)
	}
}

func (s *ResourceMemberDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitResourceMemberDecl(s)
	}
}

func (p *ModelParser) ResourceMemberDecl() (localctx IResourceMemberDeclContext) {
	this := p
	_ = this

	localctx = NewResourceMemberDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ModelParserRULE_resourceMemberDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(195)
			p.MethodDecl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(196)
			p.LocatorDecl()
		}

	}

	return localctx
}

// IMethodDeclContext is an interface to support dynamic dispatch.
type IMethodDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_annotation returns the _annotation rule contexts.
	Get_annotation() IAnnotationContext

	// GetName returns the name rule contexts.
	GetName() IIdentifierContext

	// Get_methodMemberDecl returns the _methodMemberDecl rule contexts.
	Get_methodMemberDecl() IMethodMemberDeclContext

	// Set_annotation sets the _annotation rule contexts.
	Set_annotation(IAnnotationContext)

	// SetName sets the name rule contexts.
	SetName(IIdentifierContext)

	// Set_methodMemberDecl sets the _methodMemberDecl rule contexts.
	Set_methodMemberDecl(IMethodMemberDeclContext)

	// GetAnnotations returns the annotations rule context list.
	GetAnnotations() []IAnnotationContext

	// GetMembers returns the members rule context list.
	GetMembers() []IMethodMemberDeclContext

	// SetAnnotations sets the annotations rule context list.
	SetAnnotations([]IAnnotationContext)

	// SetMembers sets the members rule context list.
	SetMembers([]IMethodMemberDeclContext)

	// GetResult returns the result attribute.
	GetResult() *concepts.Method

	// SetResult sets the result attribute.
	SetResult(*concepts.Method)

	// IsMethodDeclContext differentiates from other interfaces.
	IsMethodDeclContext()
}

type MethodDeclContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	result            *concepts.Method
	_annotation       IAnnotationContext
	annotations       []IAnnotationContext
	name              IIdentifierContext
	_methodMemberDecl IMethodMemberDeclContext
	members           []IMethodMemberDeclContext
}

func NewEmptyMethodDeclContext() *MethodDeclContext {
	var p = new(MethodDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_methodDecl
	return p
}

func (*MethodDeclContext) IsMethodDeclContext() {}

func NewMethodDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodDeclContext {
	var p = new(MethodDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_methodDecl

	return p
}

func (s *MethodDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodDeclContext) Get_annotation() IAnnotationContext { return s._annotation }

func (s *MethodDeclContext) GetName() IIdentifierContext { return s.name }

func (s *MethodDeclContext) Get_methodMemberDecl() IMethodMemberDeclContext {
	return s._methodMemberDecl
}

func (s *MethodDeclContext) Set_annotation(v IAnnotationContext) { s._annotation = v }

func (s *MethodDeclContext) SetName(v IIdentifierContext) { s.name = v }

func (s *MethodDeclContext) Set_methodMemberDecl(v IMethodMemberDeclContext) { s._methodMemberDecl = v }

func (s *MethodDeclContext) GetAnnotations() []IAnnotationContext { return s.annotations }

func (s *MethodDeclContext) GetMembers() []IMethodMemberDeclContext { return s.members }

func (s *MethodDeclContext) SetAnnotations(v []IAnnotationContext) { s.annotations = v }

func (s *MethodDeclContext) SetMembers(v []IMethodMemberDeclContext) { s.members = v }

func (s *MethodDeclContext) GetResult() *concepts.Method { return s.result }

func (s *MethodDeclContext) SetResult(v *concepts.Method) { s.result = v }

func (s *MethodDeclContext) LEFT_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserLEFT_CURLY_BRACKET, 0)
}

func (s *MethodDeclContext) RIGHT_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserRIGHT_CURLY_BRACKET, 0)
}

func (s *MethodDeclContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MethodDeclContext) METHOD() antlr.TerminalNode {
	return s.GetToken(ModelParserMETHOD, 0)
}

func (s *MethodDeclContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *MethodDeclContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *MethodDeclContext) AllMethodMemberDecl() []IMethodMemberDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodMemberDeclContext)(nil)).Elem())
	var tst = make([]IMethodMemberDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodMemberDeclContext)
		}
	}

	return tst
}

func (s *MethodDeclContext) MethodMemberDecl(i int) IMethodMemberDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodMemberDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodMemberDeclContext)
}

func (s *MethodDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterMethodDecl(s)
	}
}

func (s *MethodDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitMethodDecl(s)
	}
}

func (p *ModelParser) MethodDecl() (localctx IMethodDeclContext) {
	this := p
	_ = this

	localctx = NewMethodDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ModelParserRULE_methodDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ModelParserAT_SIGN {
		{
			p.SetState(199)

			var _x = p.Annotation()

			localctx.(*MethodDeclContext)._annotation = _x
		}
		localctx.(*MethodDeclContext).annotations = append(localctx.(*MethodDeclContext).annotations, localctx.(*MethodDeclContext)._annotation)

		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ModelParserMETHOD {
		{
			p.SetState(205)
			p.Match(ModelParserMETHOD)
		}

	}
	{
		p.SetState(208)

		var _x = p.Identifier()

		localctx.(*MethodDeclContext).name = _x
	}
	{
		p.SetState(209)
		p.Match(ModelParserLEFT_CURLY_BRACKET)
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ModelParserIN)|(1<<ModelParserOUT)|(1<<ModelParserPARAMETER)|(1<<ModelParserAT_SIGN)|(1<<ModelParserIDENTIFIER))) != 0 {
		{
			p.SetState(210)

			var _x = p.MethodMemberDecl()

			localctx.(*MethodDeclContext)._methodMemberDecl = _x
		}
		localctx.(*MethodDeclContext).members = append(localctx.(*MethodDeclContext).members, localctx.(*MethodDeclContext)._methodMemberDecl)

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(216)
		p.Match(ModelParserRIGHT_CURLY_BRACKET)
	}

	return localctx
}

// IMethodMemberDeclContext is an interface to support dynamic dispatch.
type IMethodMemberDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetResult returns the result attribute.
	GetResult() interface{}

	// SetResult sets the result attribute.
	SetResult(interface{})

	// IsMethodMemberDeclContext differentiates from other interfaces.
	IsMethodMemberDeclContext()
}

type MethodMemberDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	result interface{}
}

func NewEmptyMethodMemberDeclContext() *MethodMemberDeclContext {
	var p = new(MethodMemberDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_methodMemberDecl
	return p
}

func (*MethodMemberDeclContext) IsMethodMemberDeclContext() {}

func NewMethodMemberDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodMemberDeclContext {
	var p = new(MethodMemberDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_methodMemberDecl

	return p
}

func (s *MethodMemberDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodMemberDeclContext) GetResult() interface{} { return s.result }

func (s *MethodMemberDeclContext) SetResult(v interface{}) { s.result = v }

func (s *MethodMemberDeclContext) MethodParameterDecl() IMethodParameterDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodParameterDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodParameterDeclContext)
}

func (s *MethodMemberDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodMemberDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodMemberDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterMethodMemberDecl(s)
	}
}

func (s *MethodMemberDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitMethodMemberDecl(s)
	}
}

func (p *ModelParser) MethodMemberDecl() (localctx IMethodMemberDeclContext) {
	this := p
	_ = this

	localctx = NewMethodMemberDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ModelParserRULE_methodMemberDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.MethodParameterDecl()
	}

	return localctx
}

// IMethodParameterDeclContext is an interface to support dynamic dispatch.
type IMethodParameterDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_annotation returns the _annotation rule contexts.
	Get_annotation() IAnnotationContext

	// Get_parameterDirection returns the _parameterDirection rule contexts.
	Get_parameterDirection() IParameterDirectionContext

	// GetName returns the name rule contexts.
	GetName() IIdentifierContext

	// GetReference returns the reference rule contexts.
	GetReference() ITypeReferenceContext

	// GetDflt returns the dflt rule contexts.
	GetDflt() ILiteralContext

	// Set_annotation sets the _annotation rule contexts.
	Set_annotation(IAnnotationContext)

	// Set_parameterDirection sets the _parameterDirection rule contexts.
	Set_parameterDirection(IParameterDirectionContext)

	// SetName sets the name rule contexts.
	SetName(IIdentifierContext)

	// SetReference sets the reference rule contexts.
	SetReference(ITypeReferenceContext)

	// SetDflt sets the dflt rule contexts.
	SetDflt(ILiteralContext)

	// GetAnnotations returns the annotations rule context list.
	GetAnnotations() []IAnnotationContext

	// GetDirections returns the directions rule context list.
	GetDirections() []IParameterDirectionContext

	// SetAnnotations sets the annotations rule context list.
	SetAnnotations([]IAnnotationContext)

	// SetDirections sets the directions rule context list.
	SetDirections([]IParameterDirectionContext)

	// GetResult returns the result attribute.
	GetResult() *concepts.Parameter

	// SetResult sets the result attribute.
	SetResult(*concepts.Parameter)

	// IsMethodParameterDeclContext differentiates from other interfaces.
	IsMethodParameterDeclContext()
}

type MethodParameterDeclContext struct {
	*antlr.BaseParserRuleContext
	parser              antlr.Parser
	result              *concepts.Parameter
	_annotation         IAnnotationContext
	annotations         []IAnnotationContext
	_parameterDirection IParameterDirectionContext
	directions          []IParameterDirectionContext
	name                IIdentifierContext
	reference           ITypeReferenceContext
	dflt                ILiteralContext
}

func NewEmptyMethodParameterDeclContext() *MethodParameterDeclContext {
	var p = new(MethodParameterDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_methodParameterDecl
	return p
}

func (*MethodParameterDeclContext) IsMethodParameterDeclContext() {}

func NewMethodParameterDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodParameterDeclContext {
	var p = new(MethodParameterDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_methodParameterDecl

	return p
}

func (s *MethodParameterDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodParameterDeclContext) Get_annotation() IAnnotationContext { return s._annotation }

func (s *MethodParameterDeclContext) Get_parameterDirection() IParameterDirectionContext {
	return s._parameterDirection
}

func (s *MethodParameterDeclContext) GetName() IIdentifierContext { return s.name }

func (s *MethodParameterDeclContext) GetReference() ITypeReferenceContext { return s.reference }

func (s *MethodParameterDeclContext) GetDflt() ILiteralContext { return s.dflt }

func (s *MethodParameterDeclContext) Set_annotation(v IAnnotationContext) { s._annotation = v }

func (s *MethodParameterDeclContext) Set_parameterDirection(v IParameterDirectionContext) {
	s._parameterDirection = v
}

func (s *MethodParameterDeclContext) SetName(v IIdentifierContext) { s.name = v }

func (s *MethodParameterDeclContext) SetReference(v ITypeReferenceContext) { s.reference = v }

func (s *MethodParameterDeclContext) SetDflt(v ILiteralContext) { s.dflt = v }

func (s *MethodParameterDeclContext) GetAnnotations() []IAnnotationContext { return s.annotations }

func (s *MethodParameterDeclContext) GetDirections() []IParameterDirectionContext {
	return s.directions
}

func (s *MethodParameterDeclContext) SetAnnotations(v []IAnnotationContext) { s.annotations = v }

func (s *MethodParameterDeclContext) SetDirections(v []IParameterDirectionContext) { s.directions = v }

func (s *MethodParameterDeclContext) GetResult() *concepts.Parameter { return s.result }

func (s *MethodParameterDeclContext) SetResult(v *concepts.Parameter) { s.result = v }

func (s *MethodParameterDeclContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MethodParameterDeclContext) TypeReference() ITypeReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *MethodParameterDeclContext) PARAMETER() antlr.TerminalNode {
	return s.GetToken(ModelParserPARAMETER, 0)
}

func (s *MethodParameterDeclContext) EQUALS_SIGN() antlr.TerminalNode {
	return s.GetToken(ModelParserEQUALS_SIGN, 0)
}

func (s *MethodParameterDeclContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *MethodParameterDeclContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *MethodParameterDeclContext) AllParameterDirection() []IParameterDirectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterDirectionContext)(nil)).Elem())
	var tst = make([]IParameterDirectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterDirectionContext)
		}
	}

	return tst
}

func (s *MethodParameterDeclContext) ParameterDirection(i int) IParameterDirectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterDirectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterDirectionContext)
}

func (s *MethodParameterDeclContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *MethodParameterDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodParameterDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodParameterDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterMethodParameterDecl(s)
	}
}

func (s *MethodParameterDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitMethodParameterDecl(s)
	}
}

func (p *ModelParser) MethodParameterDecl() (localctx IMethodParameterDeclContext) {
	this := p
	_ = this

	localctx = NewMethodParameterDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ModelParserRULE_methodParameterDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ModelParserAT_SIGN {
		{
			p.SetState(220)

			var _x = p.Annotation()

			localctx.(*MethodParameterDeclContext)._annotation = _x
		}
		localctx.(*MethodParameterDeclContext).annotations = append(localctx.(*MethodParameterDeclContext).annotations, localctx.(*MethodParameterDeclContext)._annotation)

		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ModelParserPARAMETER {
		{
			p.SetState(226)
			p.Match(ModelParserPARAMETER)
		}

	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ModelParserIN || _la == ModelParserOUT {
		{
			p.SetState(229)

			var _x = p.ParameterDirection()

			localctx.(*MethodParameterDeclContext)._parameterDirection = _x
		}
		localctx.(*MethodParameterDeclContext).directions = append(localctx.(*MethodParameterDeclContext).directions, localctx.(*MethodParameterDeclContext)._parameterDirection)

		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(235)

		var _x = p.Identifier()

		localctx.(*MethodParameterDeclContext).name = _x
	}
	{
		p.SetState(236)

		var _x = p.TypeReference()

		localctx.(*MethodParameterDeclContext).reference = _x
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ModelParserEQUALS_SIGN {
		{
			p.SetState(237)
			p.Match(ModelParserEQUALS_SIGN)
		}
		{
			p.SetState(238)

			var _x = p.Literal()

			localctx.(*MethodParameterDeclContext).dflt = _x
		}

	}

	return localctx
}

// IParameterDirectionContext is an interface to support dynamic dispatch.
type IParameterDirectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIn returns the in token.
	GetIn() antlr.Token

	// GetOut returns the out token.
	GetOut() antlr.Token

	// SetIn sets the in token.
	SetIn(antlr.Token)

	// SetOut sets the out token.
	SetOut(antlr.Token)

	// IsParameterDirectionContext differentiates from other interfaces.
	IsParameterDirectionContext()
}

type ParameterDirectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	in     antlr.Token
	out    antlr.Token
}

func NewEmptyParameterDirectionContext() *ParameterDirectionContext {
	var p = new(ParameterDirectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_parameterDirection
	return p
}

func (*ParameterDirectionContext) IsParameterDirectionContext() {}

func NewParameterDirectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterDirectionContext {
	var p = new(ParameterDirectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_parameterDirection

	return p
}

func (s *ParameterDirectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterDirectionContext) GetIn() antlr.Token { return s.in }

func (s *ParameterDirectionContext) GetOut() antlr.Token { return s.out }

func (s *ParameterDirectionContext) SetIn(v antlr.Token) { s.in = v }

func (s *ParameterDirectionContext) SetOut(v antlr.Token) { s.out = v }

func (s *ParameterDirectionContext) IN() antlr.TerminalNode {
	return s.GetToken(ModelParserIN, 0)
}

func (s *ParameterDirectionContext) OUT() antlr.TerminalNode {
	return s.GetToken(ModelParserOUT, 0)
}

func (s *ParameterDirectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterDirectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterDirectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterParameterDirection(s)
	}
}

func (s *ParameterDirectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitParameterDirection(s)
	}
}

func (p *ModelParser) ParameterDirection() (localctx IParameterDirectionContext) {
	this := p
	_ = this

	localctx = NewParameterDirectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ModelParserRULE_parameterDirection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(243)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ModelParserIN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)

			var _m = p.Match(ModelParserIN)

			localctx.(*ParameterDirectionContext).in = _m
		}

	case ModelParserOUT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)

			var _m = p.Match(ModelParserOUT)

			localctx.(*ParameterDirectionContext).out = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILocatorDeclContext is an interface to support dynamic dispatch.
type ILocatorDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_annotation returns the _annotation rule contexts.
	Get_annotation() IAnnotationContext

	// GetName returns the name rule contexts.
	GetName() IIdentifierContext

	// Get_locatorMemberDecl returns the _locatorMemberDecl rule contexts.
	Get_locatorMemberDecl() ILocatorMemberDeclContext

	// Set_annotation sets the _annotation rule contexts.
	Set_annotation(IAnnotationContext)

	// SetName sets the name rule contexts.
	SetName(IIdentifierContext)

	// Set_locatorMemberDecl sets the _locatorMemberDecl rule contexts.
	Set_locatorMemberDecl(ILocatorMemberDeclContext)

	// GetAnnotations returns the annotations rule context list.
	GetAnnotations() []IAnnotationContext

	// GetMembers returns the members rule context list.
	GetMembers() []ILocatorMemberDeclContext

	// SetAnnotations sets the annotations rule context list.
	SetAnnotations([]IAnnotationContext)

	// SetMembers sets the members rule context list.
	SetMembers([]ILocatorMemberDeclContext)

	// GetResult returns the result attribute.
	GetResult() *concepts.Locator

	// SetResult sets the result attribute.
	SetResult(*concepts.Locator)

	// IsLocatorDeclContext differentiates from other interfaces.
	IsLocatorDeclContext()
}

type LocatorDeclContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	result             *concepts.Locator
	_annotation        IAnnotationContext
	annotations        []IAnnotationContext
	name               IIdentifierContext
	_locatorMemberDecl ILocatorMemberDeclContext
	members            []ILocatorMemberDeclContext
}

func NewEmptyLocatorDeclContext() *LocatorDeclContext {
	var p = new(LocatorDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_locatorDecl
	return p
}

func (*LocatorDeclContext) IsLocatorDeclContext() {}

func NewLocatorDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocatorDeclContext {
	var p = new(LocatorDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_locatorDecl

	return p
}

func (s *LocatorDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *LocatorDeclContext) Get_annotation() IAnnotationContext { return s._annotation }

func (s *LocatorDeclContext) GetName() IIdentifierContext { return s.name }

func (s *LocatorDeclContext) Get_locatorMemberDecl() ILocatorMemberDeclContext {
	return s._locatorMemberDecl
}

func (s *LocatorDeclContext) Set_annotation(v IAnnotationContext) { s._annotation = v }

func (s *LocatorDeclContext) SetName(v IIdentifierContext) { s.name = v }

func (s *LocatorDeclContext) Set_locatorMemberDecl(v ILocatorMemberDeclContext) {
	s._locatorMemberDecl = v
}

func (s *LocatorDeclContext) GetAnnotations() []IAnnotationContext { return s.annotations }

func (s *LocatorDeclContext) GetMembers() []ILocatorMemberDeclContext { return s.members }

func (s *LocatorDeclContext) SetAnnotations(v []IAnnotationContext) { s.annotations = v }

func (s *LocatorDeclContext) SetMembers(v []ILocatorMemberDeclContext) { s.members = v }

func (s *LocatorDeclContext) GetResult() *concepts.Locator { return s.result }

func (s *LocatorDeclContext) SetResult(v *concepts.Locator) { s.result = v }

func (s *LocatorDeclContext) LOCATOR() antlr.TerminalNode {
	return s.GetToken(ModelParserLOCATOR, 0)
}

func (s *LocatorDeclContext) LEFT_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserLEFT_CURLY_BRACKET, 0)
}

func (s *LocatorDeclContext) RIGHT_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserRIGHT_CURLY_BRACKET, 0)
}

func (s *LocatorDeclContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *LocatorDeclContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *LocatorDeclContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *LocatorDeclContext) AllLocatorMemberDecl() []ILocatorMemberDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILocatorMemberDeclContext)(nil)).Elem())
	var tst = make([]ILocatorMemberDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILocatorMemberDeclContext)
		}
	}

	return tst
}

func (s *LocatorDeclContext) LocatorMemberDecl(i int) ILocatorMemberDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocatorMemberDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILocatorMemberDeclContext)
}

func (s *LocatorDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocatorDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocatorDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterLocatorDecl(s)
	}
}

func (s *LocatorDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitLocatorDecl(s)
	}
}

func (p *ModelParser) LocatorDecl() (localctx ILocatorDeclContext) {
	this := p
	_ = this

	localctx = NewLocatorDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ModelParserRULE_locatorDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ModelParserAT_SIGN {
		{
			p.SetState(245)

			var _x = p.Annotation()

			localctx.(*LocatorDeclContext)._annotation = _x
		}
		localctx.(*LocatorDeclContext).annotations = append(localctx.(*LocatorDeclContext).annotations, localctx.(*LocatorDeclContext)._annotation)

		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(251)
		p.Match(ModelParserLOCATOR)
	}
	{
		p.SetState(252)

		var _x = p.Identifier()

		localctx.(*LocatorDeclContext).name = _x
	}
	{
		p.SetState(253)
		p.Match(ModelParserLEFT_CURLY_BRACKET)
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ModelParserTARGET || _la == ModelParserVARIABLE {
		{
			p.SetState(254)

			var _x = p.LocatorMemberDecl()

			localctx.(*LocatorDeclContext)._locatorMemberDecl = _x
		}
		localctx.(*LocatorDeclContext).members = append(localctx.(*LocatorDeclContext).members, localctx.(*LocatorDeclContext)._locatorMemberDecl)

		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(260)
		p.Match(ModelParserRIGHT_CURLY_BRACKET)
	}

	return localctx
}

// ILocatorMemberDeclContext is an interface to support dynamic dispatch.
type ILocatorMemberDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetResult returns the result attribute.
	GetResult() interface{}

	// SetResult sets the result attribute.
	SetResult(interface{})

	// IsLocatorMemberDeclContext differentiates from other interfaces.
	IsLocatorMemberDeclContext()
}

type LocatorMemberDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	result interface{}
}

func NewEmptyLocatorMemberDeclContext() *LocatorMemberDeclContext {
	var p = new(LocatorMemberDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_locatorMemberDecl
	return p
}

func (*LocatorMemberDeclContext) IsLocatorMemberDeclContext() {}

func NewLocatorMemberDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocatorMemberDeclContext {
	var p = new(LocatorMemberDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_locatorMemberDecl

	return p
}

func (s *LocatorMemberDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *LocatorMemberDeclContext) GetResult() interface{} { return s.result }

func (s *LocatorMemberDeclContext) SetResult(v interface{}) { s.result = v }

func (s *LocatorMemberDeclContext) LocatorTargetDecl() ILocatorTargetDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocatorTargetDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocatorTargetDeclContext)
}

func (s *LocatorMemberDeclContext) LocatorVariableDecl() ILocatorVariableDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocatorVariableDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocatorVariableDeclContext)
}

func (s *LocatorMemberDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocatorMemberDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocatorMemberDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterLocatorMemberDecl(s)
	}
}

func (s *LocatorMemberDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitLocatorMemberDecl(s)
	}
}

func (p *ModelParser) LocatorMemberDecl() (localctx ILocatorMemberDeclContext) {
	this := p
	_ = this

	localctx = NewLocatorMemberDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ModelParserRULE_locatorMemberDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(264)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ModelParserTARGET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(262)
			p.LocatorTargetDecl()
		}

	case ModelParserVARIABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(263)
			p.LocatorVariableDecl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILocatorTargetDeclContext is an interface to support dynamic dispatch.
type ILocatorTargetDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReference returns the reference rule contexts.
	GetReference() IResourceReferenceContext

	// SetReference sets the reference rule contexts.
	SetReference(IResourceReferenceContext)

	// GetResult returns the result attribute.
	GetResult() *concepts.Resource

	// SetResult sets the result attribute.
	SetResult(*concepts.Resource)

	// IsLocatorTargetDeclContext differentiates from other interfaces.
	IsLocatorTargetDeclContext()
}

type LocatorTargetDeclContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	result    *concepts.Resource
	reference IResourceReferenceContext
}

func NewEmptyLocatorTargetDeclContext() *LocatorTargetDeclContext {
	var p = new(LocatorTargetDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_locatorTargetDecl
	return p
}

func (*LocatorTargetDeclContext) IsLocatorTargetDeclContext() {}

func NewLocatorTargetDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocatorTargetDeclContext {
	var p = new(LocatorTargetDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_locatorTargetDecl

	return p
}

func (s *LocatorTargetDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *LocatorTargetDeclContext) GetReference() IResourceReferenceContext { return s.reference }

func (s *LocatorTargetDeclContext) SetReference(v IResourceReferenceContext) { s.reference = v }

func (s *LocatorTargetDeclContext) GetResult() *concepts.Resource { return s.result }

func (s *LocatorTargetDeclContext) SetResult(v *concepts.Resource) { s.result = v }

func (s *LocatorTargetDeclContext) TARGET() antlr.TerminalNode {
	return s.GetToken(ModelParserTARGET, 0)
}

func (s *LocatorTargetDeclContext) ResourceReference() IResourceReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResourceReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IResourceReferenceContext)
}

func (s *LocatorTargetDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocatorTargetDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocatorTargetDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterLocatorTargetDecl(s)
	}
}

func (s *LocatorTargetDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitLocatorTargetDecl(s)
	}
}

func (p *ModelParser) LocatorTargetDecl() (localctx ILocatorTargetDeclContext) {
	this := p
	_ = this

	localctx = NewLocatorTargetDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ModelParserRULE_locatorTargetDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(ModelParserTARGET)
	}
	{
		p.SetState(267)

		var _x = p.ResourceReference()

		localctx.(*LocatorTargetDeclContext).reference = _x
	}

	return localctx
}

// ILocatorVariableDeclContext is an interface to support dynamic dispatch.
type ILocatorVariableDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IIdentifierContext

	// SetName sets the name rule contexts.
	SetName(IIdentifierContext)

	// GetResult returns the result attribute.
	GetResult() *names.Name

	// SetResult sets the result attribute.
	SetResult(*names.Name)

	// IsLocatorVariableDeclContext differentiates from other interfaces.
	IsLocatorVariableDeclContext()
}

type LocatorVariableDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	result *names.Name
	name   IIdentifierContext
}

func NewEmptyLocatorVariableDeclContext() *LocatorVariableDeclContext {
	var p = new(LocatorVariableDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_locatorVariableDecl
	return p
}

func (*LocatorVariableDeclContext) IsLocatorVariableDeclContext() {}

func NewLocatorVariableDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocatorVariableDeclContext {
	var p = new(LocatorVariableDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_locatorVariableDecl

	return p
}

func (s *LocatorVariableDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *LocatorVariableDeclContext) GetName() IIdentifierContext { return s.name }

func (s *LocatorVariableDeclContext) SetName(v IIdentifierContext) { s.name = v }

func (s *LocatorVariableDeclContext) GetResult() *names.Name { return s.result }

func (s *LocatorVariableDeclContext) SetResult(v *names.Name) { s.result = v }

func (s *LocatorVariableDeclContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(ModelParserVARIABLE, 0)
}

func (s *LocatorVariableDeclContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *LocatorVariableDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocatorVariableDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocatorVariableDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterLocatorVariableDecl(s)
	}
}

func (s *LocatorVariableDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitLocatorVariableDecl(s)
	}
}

func (p *ModelParser) LocatorVariableDecl() (localctx ILocatorVariableDeclContext) {
	this := p
	_ = this

	localctx = NewLocatorVariableDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ModelParserRULE_locatorVariableDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.Match(ModelParserVARIABLE)
	}
	{
		p.SetState(270)

		var _x = p.Identifier()

		localctx.(*LocatorVariableDeclContext).name = _x
	}

	return localctx
}

// IResourceReferenceContext is an interface to support dynamic dispatch.
type IResourceReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IIdentifierContext

	// SetName sets the name rule contexts.
	SetName(IIdentifierContext)

	// GetResult returns the result attribute.
	GetResult() *concepts.Resource

	// SetResult sets the result attribute.
	SetResult(*concepts.Resource)

	// IsResourceReferenceContext differentiates from other interfaces.
	IsResourceReferenceContext()
}

type ResourceReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	result *concepts.Resource
	name   IIdentifierContext
}

func NewEmptyResourceReferenceContext() *ResourceReferenceContext {
	var p = new(ResourceReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_resourceReference
	return p
}

func (*ResourceReferenceContext) IsResourceReferenceContext() {}

func NewResourceReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourceReferenceContext {
	var p = new(ResourceReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_resourceReference

	return p
}

func (s *ResourceReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourceReferenceContext) GetName() IIdentifierContext { return s.name }

func (s *ResourceReferenceContext) SetName(v IIdentifierContext) { s.name = v }

func (s *ResourceReferenceContext) GetResult() *concepts.Resource { return s.result }

func (s *ResourceReferenceContext) SetResult(v *concepts.Resource) { s.result = v }

func (s *ResourceReferenceContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ResourceReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourceReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourceReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterResourceReference(s)
	}
}

func (s *ResourceReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitResourceReference(s)
	}
}

func (p *ModelParser) ResourceReference() (localctx IResourceReferenceContext) {
	this := p
	_ = this

	localctx = NewResourceReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ModelParserRULE_resourceReference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)

		var _x = p.Identifier()

		localctx.(*ResourceReferenceContext).name = _x
	}

	return localctx
}

// IErrorDeclContext is an interface to support dynamic dispatch.
type IErrorDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_annotation returns the _annotation rule contexts.
	Get_annotation() IAnnotationContext

	// GetName returns the name rule contexts.
	GetName() IIdentifierContext

	// Get_errorMemberDecl returns the _errorMemberDecl rule contexts.
	Get_errorMemberDecl() IErrorMemberDeclContext

	// Set_annotation sets the _annotation rule contexts.
	Set_annotation(IAnnotationContext)

	// SetName sets the name rule contexts.
	SetName(IIdentifierContext)

	// Set_errorMemberDecl sets the _errorMemberDecl rule contexts.
	Set_errorMemberDecl(IErrorMemberDeclContext)

	// GetAnnotations returns the annotations rule context list.
	GetAnnotations() []IAnnotationContext

	// GetMembers returns the members rule context list.
	GetMembers() []IErrorMemberDeclContext

	// SetAnnotations sets the annotations rule context list.
	SetAnnotations([]IAnnotationContext)

	// SetMembers sets the members rule context list.
	SetMembers([]IErrorMemberDeclContext)

	// GetResult returns the result attribute.
	GetResult() *concepts.Error

	// SetResult sets the result attribute.
	SetResult(*concepts.Error)

	// IsErrorDeclContext differentiates from other interfaces.
	IsErrorDeclContext()
}

type ErrorDeclContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	result           *concepts.Error
	_annotation      IAnnotationContext
	annotations      []IAnnotationContext
	name             IIdentifierContext
	_errorMemberDecl IErrorMemberDeclContext
	members          []IErrorMemberDeclContext
}

func NewEmptyErrorDeclContext() *ErrorDeclContext {
	var p = new(ErrorDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_errorDecl
	return p
}

func (*ErrorDeclContext) IsErrorDeclContext() {}

func NewErrorDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ErrorDeclContext {
	var p = new(ErrorDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_errorDecl

	return p
}

func (s *ErrorDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ErrorDeclContext) Get_annotation() IAnnotationContext { return s._annotation }

func (s *ErrorDeclContext) GetName() IIdentifierContext { return s.name }

func (s *ErrorDeclContext) Get_errorMemberDecl() IErrorMemberDeclContext { return s._errorMemberDecl }

func (s *ErrorDeclContext) Set_annotation(v IAnnotationContext) { s._annotation = v }

func (s *ErrorDeclContext) SetName(v IIdentifierContext) { s.name = v }

func (s *ErrorDeclContext) Set_errorMemberDecl(v IErrorMemberDeclContext) { s._errorMemberDecl = v }

func (s *ErrorDeclContext) GetAnnotations() []IAnnotationContext { return s.annotations }

func (s *ErrorDeclContext) GetMembers() []IErrorMemberDeclContext { return s.members }

func (s *ErrorDeclContext) SetAnnotations(v []IAnnotationContext) { s.annotations = v }

func (s *ErrorDeclContext) SetMembers(v []IErrorMemberDeclContext) { s.members = v }

func (s *ErrorDeclContext) GetResult() *concepts.Error { return s.result }

func (s *ErrorDeclContext) SetResult(v *concepts.Error) { s.result = v }

func (s *ErrorDeclContext) ERROR() antlr.TerminalNode {
	return s.GetToken(ModelParserERROR, 0)
}

func (s *ErrorDeclContext) LEFT_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserLEFT_CURLY_BRACKET, 0)
}

func (s *ErrorDeclContext) RIGHT_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(ModelParserRIGHT_CURLY_BRACKET, 0)
}

func (s *ErrorDeclContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ErrorDeclContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *ErrorDeclContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *ErrorDeclContext) AllErrorMemberDecl() []IErrorMemberDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IErrorMemberDeclContext)(nil)).Elem())
	var tst = make([]IErrorMemberDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IErrorMemberDeclContext)
		}
	}

	return tst
}

func (s *ErrorDeclContext) ErrorMemberDecl(i int) IErrorMemberDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IErrorMemberDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IErrorMemberDeclContext)
}

func (s *ErrorDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrorDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ErrorDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterErrorDecl(s)
	}
}

func (s *ErrorDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitErrorDecl(s)
	}
}

func (p *ModelParser) ErrorDecl() (localctx IErrorDeclContext) {
	this := p
	_ = this

	localctx = NewErrorDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ModelParserRULE_errorDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ModelParserAT_SIGN {
		{
			p.SetState(274)

			var _x = p.Annotation()

			localctx.(*ErrorDeclContext)._annotation = _x
		}
		localctx.(*ErrorDeclContext).annotations = append(localctx.(*ErrorDeclContext).annotations, localctx.(*ErrorDeclContext)._annotation)

		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(280)
		p.Match(ModelParserERROR)
	}
	{
		p.SetState(281)

		var _x = p.Identifier()

		localctx.(*ErrorDeclContext).name = _x
	}
	{
		p.SetState(282)
		p.Match(ModelParserLEFT_CURLY_BRACKET)
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ModelParserCODE {
		{
			p.SetState(283)

			var _x = p.ErrorMemberDecl()

			localctx.(*ErrorDeclContext)._errorMemberDecl = _x
		}
		localctx.(*ErrorDeclContext).members = append(localctx.(*ErrorDeclContext).members, localctx.(*ErrorDeclContext)._errorMemberDecl)

		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(289)
		p.Match(ModelParserRIGHT_CURLY_BRACKET)
	}

	return localctx
}

// IErrorMemberDeclContext is an interface to support dynamic dispatch.
type IErrorMemberDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetResult returns the result attribute.
	GetResult() interface{}

	// SetResult sets the result attribute.
	SetResult(interface{})

	// IsErrorMemberDeclContext differentiates from other interfaces.
	IsErrorMemberDeclContext()
}

type ErrorMemberDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	result interface{}
}

func NewEmptyErrorMemberDeclContext() *ErrorMemberDeclContext {
	var p = new(ErrorMemberDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_errorMemberDecl
	return p
}

func (*ErrorMemberDeclContext) IsErrorMemberDeclContext() {}

func NewErrorMemberDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ErrorMemberDeclContext {
	var p = new(ErrorMemberDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_errorMemberDecl

	return p
}

func (s *ErrorMemberDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ErrorMemberDeclContext) GetResult() interface{} { return s.result }

func (s *ErrorMemberDeclContext) SetResult(v interface{}) { s.result = v }

func (s *ErrorMemberDeclContext) ErrorCodeDecl() IErrorCodeDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IErrorCodeDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IErrorCodeDeclContext)
}

func (s *ErrorMemberDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrorMemberDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ErrorMemberDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterErrorMemberDecl(s)
	}
}

func (s *ErrorMemberDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitErrorMemberDecl(s)
	}
}

func (p *ModelParser) ErrorMemberDecl() (localctx IErrorMemberDeclContext) {
	this := p
	_ = this

	localctx = NewErrorMemberDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ModelParserRULE_errorMemberDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.ErrorCodeDecl()
	}

	return localctx
}

// IErrorCodeDeclContext is an interface to support dynamic dispatch.
type IErrorCodeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCode returns the code token.
	GetCode() antlr.Token

	// SetCode sets the code token.
	SetCode(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() int

	// SetResult sets the result attribute.
	SetResult(int)

	// IsErrorCodeDeclContext differentiates from other interfaces.
	IsErrorCodeDeclContext()
}

type ErrorCodeDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	result int
	code   antlr.Token
}

func NewEmptyErrorCodeDeclContext() *ErrorCodeDeclContext {
	var p = new(ErrorCodeDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_errorCodeDecl
	return p
}

func (*ErrorCodeDeclContext) IsErrorCodeDeclContext() {}

func NewErrorCodeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ErrorCodeDeclContext {
	var p = new(ErrorCodeDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_errorCodeDecl

	return p
}

func (s *ErrorCodeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ErrorCodeDeclContext) GetCode() antlr.Token { return s.code }

func (s *ErrorCodeDeclContext) SetCode(v antlr.Token) { s.code = v }

func (s *ErrorCodeDeclContext) GetResult() int { return s.result }

func (s *ErrorCodeDeclContext) SetResult(v int) { s.result = v }

func (s *ErrorCodeDeclContext) CODE() antlr.TerminalNode {
	return s.GetToken(ModelParserCODE, 0)
}

func (s *ErrorCodeDeclContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(ModelParserINTEGER_LITERAL, 0)
}

func (s *ErrorCodeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrorCodeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ErrorCodeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterErrorCodeDecl(s)
	}
}

func (s *ErrorCodeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitErrorCodeDecl(s)
	}
}

func (p *ModelParser) ErrorCodeDecl() (localctx IErrorCodeDeclContext) {
	this := p
	_ = this

	localctx = NewErrorCodeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ModelParserRULE_errorCodeDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(ModelParserCODE)
	}
	{
		p.SetState(294)

		var _m = p.Match(ModelParserINTEGER_LITERAL)

		localctx.(*ErrorCodeDeclContext).code = _m
	}

	return localctx
}

// IAnnotationContext is an interface to support dynamic dispatch.
type IAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IIdentifierContext

	// GetParameters returns the parameters rule contexts.
	GetParameters() IAnnotationParametersContext

	// SetName sets the name rule contexts.
	SetName(IIdentifierContext)

	// SetParameters sets the parameters rule contexts.
	SetParameters(IAnnotationParametersContext)

	// GetResult returns the result attribute.
	GetResult() *concepts.Annotation

	// SetResult sets the result attribute.
	SetResult(*concepts.Annotation)

	// IsAnnotationContext differentiates from other interfaces.
	IsAnnotationContext()
}

type AnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	result     *concepts.Annotation
	name       IIdentifierContext
	parameters IAnnotationParametersContext
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_annotation
	return p
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationContext) GetName() IIdentifierContext { return s.name }

func (s *AnnotationContext) GetParameters() IAnnotationParametersContext { return s.parameters }

func (s *AnnotationContext) SetName(v IIdentifierContext) { s.name = v }

func (s *AnnotationContext) SetParameters(v IAnnotationParametersContext) { s.parameters = v }

func (s *AnnotationContext) GetResult() *concepts.Annotation { return s.result }

func (s *AnnotationContext) SetResult(v *concepts.Annotation) { s.result = v }

func (s *AnnotationContext) AT_SIGN() antlr.TerminalNode {
	return s.GetToken(ModelParserAT_SIGN, 0)
}

func (s *AnnotationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AnnotationContext) AnnotationParameters() IAnnotationParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationParametersContext)
}

func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterAnnotation(s)
	}
}

func (s *AnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitAnnotation(s)
	}
}

func (p *ModelParser) Annotation() (localctx IAnnotationContext) {
	this := p
	_ = this

	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ModelParserRULE_annotation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Match(ModelParserAT_SIGN)
	}
	{
		p.SetState(297)

		var _x = p.Identifier()

		localctx.(*AnnotationContext).name = _x
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ModelParserLEFT_PARENTHESIS {
		{
			p.SetState(298)

			var _x = p.AnnotationParameters()

			localctx.(*AnnotationContext).parameters = _x
		}

	}

	return localctx
}

// IAnnotationParametersContext is an interface to support dynamic dispatch.
type IAnnotationParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_annotationParameter returns the _annotationParameter rule contexts.
	Get_annotationParameter() IAnnotationParameterContext

	// Set_annotationParameter sets the _annotationParameter rule contexts.
	Set_annotationParameter(IAnnotationParameterContext)

	// GetParameters returns the parameters rule context list.
	GetParameters() []IAnnotationParameterContext

	// SetParameters sets the parameters rule context list.
	SetParameters([]IAnnotationParameterContext)

	// GetResult returns the result attribute.
	GetResult() map[string]interface{}

	// SetResult sets the result attribute.
	SetResult(map[string]interface{})

	// IsAnnotationParametersContext differentiates from other interfaces.
	IsAnnotationParametersContext()
}

type AnnotationParametersContext struct {
	*antlr.BaseParserRuleContext
	parser               antlr.Parser
	result               map[string]interface{}
	_annotationParameter IAnnotationParameterContext
	parameters           []IAnnotationParameterContext
}

func NewEmptyAnnotationParametersContext() *AnnotationParametersContext {
	var p = new(AnnotationParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_annotationParameters
	return p
}

func (*AnnotationParametersContext) IsAnnotationParametersContext() {}

func NewAnnotationParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationParametersContext {
	var p = new(AnnotationParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_annotationParameters

	return p
}

func (s *AnnotationParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationParametersContext) Get_annotationParameter() IAnnotationParameterContext {
	return s._annotationParameter
}

func (s *AnnotationParametersContext) Set_annotationParameter(v IAnnotationParameterContext) {
	s._annotationParameter = v
}

func (s *AnnotationParametersContext) GetParameters() []IAnnotationParameterContext {
	return s.parameters
}

func (s *AnnotationParametersContext) SetParameters(v []IAnnotationParameterContext) {
	s.parameters = v
}

func (s *AnnotationParametersContext) GetResult() map[string]interface{} { return s.result }

func (s *AnnotationParametersContext) SetResult(v map[string]interface{}) { s.result = v }

func (s *AnnotationParametersContext) LEFT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(ModelParserLEFT_PARENTHESIS, 0)
}

func (s *AnnotationParametersContext) RIGHT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(ModelParserRIGHT_PARENTHESIS, 0)
}

func (s *AnnotationParametersContext) AllAnnotationParameter() []IAnnotationParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationParameterContext)(nil)).Elem())
	var tst = make([]IAnnotationParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationParameterContext)
		}
	}

	return tst
}

func (s *AnnotationParametersContext) AnnotationParameter(i int) IAnnotationParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationParameterContext)
}

func (s *AnnotationParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterAnnotationParameters(s)
	}
}

func (s *AnnotationParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitAnnotationParameters(s)
	}
}

func (p *ModelParser) AnnotationParameters() (localctx IAnnotationParametersContext) {
	this := p
	_ = this

	localctx = NewAnnotationParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ModelParserRULE_annotationParameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Match(ModelParserLEFT_PARENTHESIS)
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ModelParserIDENTIFIER {
		{
			p.SetState(302)

			var _x = p.AnnotationParameter()

			localctx.(*AnnotationParametersContext)._annotationParameter = _x
		}
		localctx.(*AnnotationParametersContext).parameters = append(localctx.(*AnnotationParametersContext).parameters, localctx.(*AnnotationParametersContext)._annotationParameter)

		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(308)
		p.Match(ModelParserRIGHT_PARENTHESIS)
	}

	return localctx
}

// IAnnotationParameterContext is an interface to support dynamic dispatch.
type IAnnotationParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name attribute.
	GetName() string

	// GetValue returns the value attribute.
	GetValue() interface{}

	// SetName sets the name attribute.
	SetName(string)

	// SetValue sets the value attribute.
	SetValue(interface{})

	// IsAnnotationParameterContext differentiates from other interfaces.
	IsAnnotationParameterContext()
}

type AnnotationParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   string
	value  interface{}
}

func NewEmptyAnnotationParameterContext() *AnnotationParameterContext {
	var p = new(AnnotationParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_annotationParameter
	return p
}

func (*AnnotationParameterContext) IsAnnotationParameterContext() {}

func NewAnnotationParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationParameterContext {
	var p = new(AnnotationParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_annotationParameter

	return p
}

func (s *AnnotationParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationParameterContext) GetName() string { return s.name }

func (s *AnnotationParameterContext) GetValue() interface{} { return s.value }

func (s *AnnotationParameterContext) SetName(v string) { s.name = v }

func (s *AnnotationParameterContext) SetValue(v interface{}) { s.value = v }

func (s *AnnotationParameterContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AnnotationParameterContext) EQUALS_SIGN() antlr.TerminalNode {
	return s.GetToken(ModelParserEQUALS_SIGN, 0)
}

func (s *AnnotationParameterContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *AnnotationParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterAnnotationParameter(s)
	}
}

func (s *AnnotationParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitAnnotationParameter(s)
	}
}

func (p *ModelParser) AnnotationParameter() (localctx IAnnotationParameterContext) {
	this := p
	_ = this

	localctx = NewAnnotationParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ModelParserRULE_annotationParameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Identifier()
	}
	{
		p.SetState(311)
		p.Match(ModelParserEQUALS_SIGN)
	}
	{
		p.SetState(312)
		p.Literal()
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetResult returns the result attribute.
	GetResult() interface{}

	// SetResult sets the result attribute.
	SetResult(interface{})

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	result interface{}
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) GetResult() interface{} { return s.result }

func (s *LiteralContext) SetResult(v interface{}) { s.result = v }

func (s *LiteralContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *LiteralContext) IntegerLiteral() IIntegerLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *LiteralContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *ModelParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ModelParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(317)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ModelParserFALSE, ModelParserTRUE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(314)
			p.BooleanLiteral()
		}

	case ModelParserINTEGER_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(315)
			p.IntegerLiteral()
		}

	case ModelParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(316)
			p.StringLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetResult returns the result attribute.
	GetResult() bool

	// SetResult sets the result attribute.
	SetResult(bool)

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	result bool
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) GetResult() bool { return s.result }

func (s *BooleanLiteralContext) SetResult(v bool) { s.result = v }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(ModelParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(ModelParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (p *ModelParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	this := p
	_ = this

	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ModelParserRULE_booleanLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(319)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ModelParserFALSE || _la == ModelParserTRUE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetResult returns the result attribute.
	GetResult() int

	// SetResult sets the result attribute.
	SetResult(int)

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	result int
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_integerLiteral
	return p
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) GetResult() int { return s.result }

func (s *IntegerLiteralContext) SetResult(v int) { s.result = v }

func (s *IntegerLiteralContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(ModelParserINTEGER_LITERAL, 0)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (p *ModelParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	this := p
	_ = this

	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ModelParserRULE_integerLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.Match(ModelParserINTEGER_LITERAL)
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetResult returns the result attribute.
	GetResult() string

	// SetResult sets the result attribute.
	SetResult(string)

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	result string
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) GetResult() string { return s.result }

func (s *StringLiteralContext) SetResult(v string) { s.result = v }

func (s *StringLiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ModelParserSTRING_LITERAL, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (p *ModelParser) StringLiteral() (localctx IStringLiteralContext) {
	this := p
	_ = this

	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ModelParserRULE_stringLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(323)
		p.Match(ModelParserSTRING_LITERAL)
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() *names.Name

	// GetText returns the text attribute.
	GetText() string

	// SetResult sets the result attribute.
	SetResult(*names.Name)

	// SetText sets the text attribute.
	SetText(string)

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	result *names.Name
	text   string
	id     antlr.Token
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) GetId() antlr.Token { return s.id }

func (s *IdentifierContext) SetId(v antlr.Token) { s.id = v }

func (s *IdentifierContext) GetResult() *names.Name { return s.result }

func (s *IdentifierContext) GetText() string { return s.text }

func (s *IdentifierContext) SetResult(v *names.Name) { s.result = v }

func (s *IdentifierContext) SetText(v string) { s.text = v }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ModelParserIDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *ModelParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ModelParserRULE_identifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)

		var _m = p.Match(ModelParserIDENTIFIER)

		localctx.(*IdentifierContext).id = _m
	}

	return localctx
}
