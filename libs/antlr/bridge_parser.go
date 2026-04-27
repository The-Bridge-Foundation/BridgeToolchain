// Code generated from antlr/Bridge.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Bridge

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 88, 487,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 3, 2, 6, 2, 108, 10,
	2, 13, 2, 14, 2, 109, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 7, 3, 117, 10, 3, 12,
	3, 14, 3, 120, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 126, 10, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 134, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	7, 7, 7, 141, 10, 7, 12, 7, 14, 7, 144, 11, 7, 3, 8, 3, 8, 5, 8, 148, 10,
	8, 3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 154, 10, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 5, 10, 182, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 6, 10, 188, 10,
	10, 13, 10, 14, 10, 189, 7, 10, 192, 10, 10, 12, 10, 14, 10, 195, 11, 10,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 215, 10, 16,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 5, 19, 224, 10, 19, 3,
	19, 5, 19, 227, 10, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 5, 20, 234,
	10, 20, 3, 21, 3, 21, 3, 21, 7, 21, 239, 10, 21, 12, 21, 14, 21, 242, 11,
	21, 3, 22, 3, 22, 5, 22, 246, 10, 22, 3, 22, 3, 22, 5, 22, 250, 10, 22,
	3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 261,
	10, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 7, 27, 274, 10, 27, 12, 27, 14, 27, 277, 11, 27, 3, 28, 3,
	28, 3, 28, 3, 28, 5, 28, 283, 10, 28, 3, 29, 3, 29, 3, 29, 5, 29, 288,
	10, 29, 3, 29, 3, 29, 3, 29, 5, 29, 293, 10, 29, 6, 29, 295, 10, 29, 13,
	29, 14, 29, 296, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 5, 31, 308, 10, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 323, 10, 34, 3, 35,
	3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 5, 37, 334, 10,
	37, 3, 37, 3, 37, 7, 37, 338, 10, 37, 12, 37, 14, 37, 341, 11, 37, 3, 37,
	5, 37, 344, 10, 37, 3, 38, 3, 38, 7, 38, 348, 10, 38, 12, 38, 14, 38, 351,
	11, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40,
	3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 5, 40, 369, 10, 40, 3,
	41, 3, 41, 7, 41, 373, 10, 41, 12, 41, 14, 41, 376, 11, 41, 3, 41, 3, 41,
	5, 41, 380, 10, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 7, 42, 388,
	10, 42, 12, 42, 14, 42, 391, 11, 42, 5, 42, 393, 10, 42, 3, 42, 3, 42,
	5, 42, 397, 10, 42, 5, 42, 399, 10, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3,
	43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 5, 43,
	415, 10, 43, 3, 44, 3, 44, 5, 44, 419, 10, 44, 3, 44, 3, 44, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 46, 3, 46, 5, 46, 429, 10, 46, 3, 46, 3, 46, 5, 46,
	433, 10, 46, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 7,
	48, 443, 10, 48, 12, 48, 14, 48, 446, 11, 48, 5, 48, 448, 10, 48, 3, 49,
	3, 49, 5, 49, 452, 10, 49, 3, 49, 3, 49, 5, 49, 456, 10, 49, 3, 49, 3,
	49, 5, 49, 460, 10, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51,
	5, 51, 469, 10, 51, 3, 51, 5, 51, 472, 10, 51, 3, 52, 3, 52, 7, 52, 476,
	10, 52, 12, 52, 14, 52, 479, 11, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53,
	3, 53, 3, 53, 2, 3, 18, 54, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
	62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96,
	98, 100, 102, 104, 2, 9, 5, 2, 65, 66, 69, 75, 77, 79, 3, 2, 67, 68, 4,
	2, 55, 55, 80, 84, 4, 2, 56, 57, 61, 64, 4, 2, 34, 34, 36, 37, 4, 2, 39,
	39, 88, 88, 3, 2, 3, 5, 2, 528, 2, 107, 3, 2, 2, 2, 4, 113, 3, 2, 2, 2,
	6, 121, 3, 2, 2, 2, 8, 133, 3, 2, 2, 2, 10, 135, 3, 2, 2, 2, 12, 137, 3,
	2, 2, 2, 14, 145, 3, 2, 2, 2, 16, 153, 3, 2, 2, 2, 18, 181, 3, 2, 2, 2,
	20, 196, 3, 2, 2, 2, 22, 199, 3, 2, 2, 2, 24, 202, 3, 2, 2, 2, 26, 204,
	3, 2, 2, 2, 28, 207, 3, 2, 2, 2, 30, 214, 3, 2, 2, 2, 32, 216, 3, 2, 2,
	2, 34, 218, 3, 2, 2, 2, 36, 223, 3, 2, 2, 2, 38, 230, 3, 2, 2, 2, 40, 235,
	3, 2, 2, 2, 42, 245, 3, 2, 2, 2, 44, 253, 3, 2, 2, 2, 46, 260, 3, 2, 2,
	2, 48, 262, 3, 2, 2, 2, 50, 265, 3, 2, 2, 2, 52, 268, 3, 2, 2, 2, 54, 278,
	3, 2, 2, 2, 56, 287, 3, 2, 2, 2, 58, 298, 3, 2, 2, 2, 60, 302, 3, 2, 2,
	2, 62, 311, 3, 2, 2, 2, 64, 315, 3, 2, 2, 2, 66, 322, 3, 2, 2, 2, 68, 324,
	3, 2, 2, 2, 70, 326, 3, 2, 2, 2, 72, 333, 3, 2, 2, 2, 74, 345, 3, 2, 2,
	2, 76, 354, 3, 2, 2, 2, 78, 368, 3, 2, 2, 2, 80, 379, 3, 2, 2, 2, 82, 398,
	3, 2, 2, 2, 84, 414, 3, 2, 2, 2, 86, 416, 3, 2, 2, 2, 88, 422, 3, 2, 2,
	2, 90, 428, 3, 2, 2, 2, 92, 434, 3, 2, 2, 2, 94, 447, 3, 2, 2, 2, 96, 449,
	3, 2, 2, 2, 98, 463, 3, 2, 2, 2, 100, 465, 3, 2, 2, 2, 102, 473, 3, 2,
	2, 2, 104, 482, 3, 2, 2, 2, 106, 108, 5, 8, 5, 2, 107, 106, 3, 2, 2, 2,
	108, 109, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110,
	111, 3, 2, 2, 2, 111, 112, 7, 2, 2, 3, 112, 3, 3, 2, 2, 2, 113, 118, 7,
	88, 2, 2, 114, 115, 7, 49, 2, 2, 115, 117, 7, 88, 2, 2, 116, 114, 3, 2,
	2, 2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2,
	119, 5, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 122, 7, 52, 2, 2, 122, 123,
	7, 88, 2, 2, 123, 125, 7, 43, 2, 2, 124, 126, 5, 4, 3, 2, 125, 124, 3,
	2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 128, 7, 44, 2,
	2, 128, 7, 3, 2, 2, 2, 129, 134, 5, 6, 4, 2, 130, 134, 5, 96, 49, 2, 131,
	134, 5, 100, 51, 2, 132, 134, 5, 104, 53, 2, 133, 129, 3, 2, 2, 2, 133,
	130, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133, 132, 3, 2, 2, 2, 134, 9, 3,
	2, 2, 2, 135, 136, 9, 2, 2, 2, 136, 11, 3, 2, 2, 2, 137, 142, 5, 18, 10,
	2, 138, 139, 7, 49, 2, 2, 139, 141, 5, 18, 10, 2, 140, 138, 3, 2, 2, 2,
	141, 144, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143,
	13, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 145, 147, 7, 47, 2, 2, 146, 148,
	5, 12, 7, 2, 147, 146, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 149, 3, 2,
	2, 2, 149, 150, 7, 48, 2, 2, 150, 15, 3, 2, 2, 2, 151, 154, 5, 56, 29,
	2, 152, 154, 7, 88, 2, 2, 153, 151, 3, 2, 2, 2, 153, 152, 3, 2, 2, 2, 154,
	155, 3, 2, 2, 2, 155, 156, 5, 88, 45, 2, 156, 17, 3, 2, 2, 2, 157, 158,
	8, 10, 1, 2, 158, 159, 7, 43, 2, 2, 159, 160, 5, 18, 10, 2, 160, 161, 7,
	44, 2, 2, 161, 182, 3, 2, 2, 2, 162, 163, 7, 43, 2, 2, 163, 164, 5, 46,
	24, 2, 164, 165, 7, 44, 2, 2, 165, 182, 3, 2, 2, 2, 166, 182, 5, 30, 16,
	2, 167, 182, 5, 42, 22, 2, 168, 182, 5, 96, 49, 2, 169, 182, 5, 16, 9,
	2, 170, 182, 5, 56, 29, 2, 171, 182, 5, 14, 8, 2, 172, 173, 7, 76, 2, 2,
	173, 182, 7, 88, 2, 2, 174, 182, 7, 20, 2, 2, 175, 182, 7, 21, 2, 2, 176,
	182, 7, 22, 2, 2, 177, 182, 7, 40, 2, 2, 178, 182, 7, 39, 2, 2, 179, 182,
	7, 41, 2, 2, 180, 182, 7, 88, 2, 2, 181, 157, 3, 2, 2, 2, 181, 162, 3,
	2, 2, 2, 181, 166, 3, 2, 2, 2, 181, 167, 3, 2, 2, 2, 181, 168, 3, 2, 2,
	2, 181, 169, 3, 2, 2, 2, 181, 170, 3, 2, 2, 2, 181, 171, 3, 2, 2, 2, 181,
	172, 3, 2, 2, 2, 181, 174, 3, 2, 2, 2, 181, 175, 3, 2, 2, 2, 181, 176,
	3, 2, 2, 2, 181, 177, 3, 2, 2, 2, 181, 178, 3, 2, 2, 2, 181, 179, 3, 2,
	2, 2, 181, 180, 3, 2, 2, 2, 182, 193, 3, 2, 2, 2, 183, 187, 12, 16, 2,
	2, 184, 185, 5, 10, 6, 2, 185, 186, 5, 18, 10, 2, 186, 188, 3, 2, 2, 2,
	187, 184, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 189,
	190, 3, 2, 2, 2, 190, 192, 3, 2, 2, 2, 191, 183, 3, 2, 2, 2, 192, 195,
	3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 19, 3, 2,
	2, 2, 195, 193, 3, 2, 2, 2, 196, 197, 7, 38, 2, 2, 197, 198, 5, 18, 10,
	2, 198, 21, 3, 2, 2, 2, 199, 200, 7, 58, 2, 2, 200, 201, 5, 18, 10, 2,
	201, 23, 3, 2, 2, 2, 202, 203, 9, 3, 2, 2, 203, 25, 3, 2, 2, 2, 204, 205,
	5, 24, 13, 2, 205, 206, 7, 88, 2, 2, 206, 27, 3, 2, 2, 2, 207, 208, 7,
	88, 2, 2, 208, 209, 5, 24, 13, 2, 209, 29, 3, 2, 2, 2, 210, 215, 5, 20,
	11, 2, 211, 215, 5, 22, 12, 2, 212, 215, 5, 26, 14, 2, 213, 215, 5, 28,
	15, 2, 214, 210, 3, 2, 2, 2, 214, 211, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2,
	214, 213, 3, 2, 2, 2, 215, 31, 3, 2, 2, 2, 216, 217, 9, 4, 2, 2, 217, 33,
	3, 2, 2, 2, 218, 219, 5, 32, 17, 2, 219, 220, 5, 18, 10, 2, 220, 35, 3,
	2, 2, 2, 221, 224, 5, 56, 29, 2, 222, 224, 7, 88, 2, 2, 223, 221, 3, 2,
	2, 2, 223, 222, 3, 2, 2, 2, 224, 226, 3, 2, 2, 2, 225, 227, 5, 88, 45,
	2, 226, 225, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228,
	229, 5, 34, 18, 2, 229, 37, 3, 2, 2, 2, 230, 233, 7, 35, 2, 2, 231, 234,
	5, 78, 40, 2, 232, 234, 5, 18, 10, 2, 233, 231, 3, 2, 2, 2, 233, 232, 3,
	2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 39, 3, 2, 2, 2, 235, 240, 5, 18, 10,
	2, 236, 237, 7, 49, 2, 2, 237, 239, 5, 18, 10, 2, 238, 236, 3, 2, 2, 2,
	239, 242, 3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241,
	41, 3, 2, 2, 2, 242, 240, 3, 2, 2, 2, 243, 246, 5, 54, 28, 2, 244, 246,
	7, 88, 2, 2, 245, 243, 3, 2, 2, 2, 245, 244, 3, 2, 2, 2, 246, 247, 3, 2,
	2, 2, 247, 249, 7, 43, 2, 2, 248, 250, 5, 40, 21, 2, 249, 248, 3, 2, 2,
	2, 249, 250, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 252, 7, 44, 2, 2, 252,
	43, 3, 2, 2, 2, 253, 254, 9, 5, 2, 2, 254, 45, 3, 2, 2, 2, 255, 261, 5,
	18, 10, 2, 256, 257, 5, 18, 10, 2, 257, 258, 5, 44, 23, 2, 258, 259, 5,
	18, 10, 2, 259, 261, 3, 2, 2, 2, 260, 255, 3, 2, 2, 2, 260, 256, 3, 2,
	2, 2, 261, 47, 3, 2, 2, 2, 262, 263, 7, 29, 2, 2, 263, 264, 5, 80, 41,
	2, 264, 49, 3, 2, 2, 2, 265, 266, 7, 29, 2, 2, 266, 267, 5, 52, 27, 2,
	267, 51, 3, 2, 2, 2, 268, 269, 7, 28, 2, 2, 269, 270, 5, 46, 24, 2, 270,
	275, 5, 80, 41, 2, 271, 274, 5, 50, 26, 2, 272, 274, 5, 48, 25, 2, 273,
	271, 3, 2, 2, 2, 273, 272, 3, 2, 2, 2, 274, 277, 3, 2, 2, 2, 275, 273,
	3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 53, 3, 2, 2, 2, 277, 275, 3, 2,
	2, 2, 278, 279, 7, 88, 2, 2, 279, 282, 7, 53, 2, 2, 280, 283, 5, 42, 22,
	2, 281, 283, 7, 88, 2, 2, 282, 280, 3, 2, 2, 2, 282, 281, 3, 2, 2, 2, 283,
	55, 3, 2, 2, 2, 284, 288, 5, 54, 28, 2, 285, 288, 5, 42, 22, 2, 286, 288,
	7, 88, 2, 2, 287, 284, 3, 2, 2, 2, 287, 285, 3, 2, 2, 2, 287, 286, 3, 2,
	2, 2, 288, 294, 3, 2, 2, 2, 289, 292, 7, 50, 2, 2, 290, 293, 5, 42, 22,
	2, 291, 293, 7, 88, 2, 2, 292, 290, 3, 2, 2, 2, 292, 291, 3, 2, 2, 2, 293,
	295, 3, 2, 2, 2, 294, 289, 3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 294,
	3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 57, 3, 2, 2, 2, 298, 299, 7, 26,
	2, 2, 299, 300, 5, 18, 10, 2, 300, 301, 5, 80, 41, 2, 301, 59, 3, 2, 2,
	2, 302, 303, 7, 27, 2, 2, 303, 307, 7, 88, 2, 2, 304, 308, 5, 56, 29, 2,
	305, 308, 7, 88, 2, 2, 306, 308, 5, 14, 8, 2, 307, 304, 3, 2, 2, 2, 307,
	305, 3, 2, 2, 2, 307, 306, 3, 2, 2, 2, 308, 309, 3, 2, 2, 2, 309, 310,
	5, 80, 41, 2, 310, 61, 3, 2, 2, 2, 311, 312, 7, 24, 2, 2, 312, 313, 5,
	46, 24, 2, 313, 314, 5, 80, 41, 2, 314, 63, 3, 2, 2, 2, 315, 316, 7, 25,
	2, 2, 316, 317, 5, 80, 41, 2, 317, 65, 3, 2, 2, 2, 318, 323, 5, 64, 33,
	2, 319, 323, 5, 62, 32, 2, 320, 323, 5, 60, 31, 2, 321, 323, 5, 58, 30,
	2, 322, 318, 3, 2, 2, 2, 322, 319, 3, 2, 2, 2, 322, 320, 3, 2, 2, 2, 322,
	321, 3, 2, 2, 2, 323, 67, 3, 2, 2, 2, 324, 325, 9, 6, 2, 2, 325, 69, 3,
	2, 2, 2, 326, 327, 7, 31, 2, 2, 327, 328, 5, 18, 10, 2, 328, 329, 5, 74,
	38, 2, 329, 71, 3, 2, 2, 2, 330, 331, 7, 32, 2, 2, 331, 334, 5, 18, 10,
	2, 332, 334, 7, 33, 2, 2, 333, 330, 3, 2, 2, 2, 333, 332, 3, 2, 2, 2, 334,
	335, 3, 2, 2, 2, 335, 339, 7, 60, 2, 2, 336, 338, 5, 78, 40, 2, 337, 336,
	3, 2, 2, 2, 338, 341, 3, 2, 2, 2, 339, 337, 3, 2, 2, 2, 339, 340, 3, 2,
	2, 2, 340, 343, 3, 2, 2, 2, 341, 339, 3, 2, 2, 2, 342, 344, 5, 68, 35,
	2, 343, 342, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344, 73, 3, 2, 2, 2, 345,
	349, 7, 45, 2, 2, 346, 348, 5, 72, 37, 2, 347, 346, 3, 2, 2, 2, 348, 351,
	3, 2, 2, 2, 349, 347, 3, 2, 2, 2, 349, 350, 3, 2, 2, 2, 350, 352, 3, 2,
	2, 2, 351, 349, 3, 2, 2, 2, 352, 353, 7, 46, 2, 2, 353, 75, 3, 2, 2, 2,
	354, 355, 7, 6, 2, 2, 355, 356, 7, 88, 2, 2, 356, 77, 3, 2, 2, 2, 357,
	369, 5, 56, 29, 2, 358, 369, 5, 42, 22, 2, 359, 369, 5, 8, 5, 2, 360, 369,
	5, 36, 19, 2, 361, 369, 5, 38, 20, 2, 362, 369, 5, 68, 35, 2, 363, 369,
	5, 66, 34, 2, 364, 369, 5, 70, 36, 2, 365, 369, 5, 52, 27, 2, 366, 369,
	5, 30, 16, 2, 367, 369, 5, 76, 39, 2, 368, 357, 3, 2, 2, 2, 368, 358, 3,
	2, 2, 2, 368, 359, 3, 2, 2, 2, 368, 360, 3, 2, 2, 2, 368, 361, 3, 2, 2,
	2, 368, 362, 3, 2, 2, 2, 368, 363, 3, 2, 2, 2, 368, 364, 3, 2, 2, 2, 368,
	365, 3, 2, 2, 2, 368, 366, 3, 2, 2, 2, 368, 367, 3, 2, 2, 2, 369, 79, 3,
	2, 2, 2, 370, 374, 7, 45, 2, 2, 371, 373, 5, 78, 40, 2, 372, 371, 3, 2,
	2, 2, 373, 376, 3, 2, 2, 2, 374, 372, 3, 2, 2, 2, 374, 375, 3, 2, 2, 2,
	375, 377, 3, 2, 2, 2, 376, 374, 3, 2, 2, 2, 377, 380, 7, 46, 2, 2, 378,
	380, 5, 78, 40, 2, 379, 370, 3, 2, 2, 2, 379, 378, 3, 2, 2, 2, 380, 81,
	3, 2, 2, 2, 381, 399, 7, 7, 2, 2, 382, 383, 7, 7, 2, 2, 383, 392, 7, 43,
	2, 2, 384, 389, 5, 90, 46, 2, 385, 386, 7, 49, 2, 2, 386, 388, 5, 90, 46,
	2, 387, 385, 3, 2, 2, 2, 388, 391, 3, 2, 2, 2, 389, 387, 3, 2, 2, 2, 389,
	390, 3, 2, 2, 2, 390, 393, 3, 2, 2, 2, 391, 389, 3, 2, 2, 2, 392, 384,
	3, 2, 2, 2, 392, 393, 3, 2, 2, 2, 393, 394, 3, 2, 2, 2, 394, 396, 7, 44,
	2, 2, 395, 397, 5, 90, 46, 2, 396, 395, 3, 2, 2, 2, 396, 397, 3, 2, 2,
	2, 397, 399, 3, 2, 2, 2, 398, 381, 3, 2, 2, 2, 398, 382, 3, 2, 2, 2, 399,
	83, 3, 2, 2, 2, 400, 415, 5, 82, 42, 2, 401, 415, 7, 23, 2, 2, 402, 415,
	7, 9, 2, 2, 403, 415, 7, 10, 2, 2, 404, 415, 7, 8, 2, 2, 405, 415, 7, 11,
	2, 2, 406, 415, 7, 12, 2, 2, 407, 415, 7, 13, 2, 2, 408, 415, 7, 14, 2,
	2, 409, 415, 7, 15, 2, 2, 410, 415, 7, 16, 2, 2, 411, 415, 7, 17, 2, 2,
	412, 415, 7, 18, 2, 2, 413, 415, 7, 19, 2, 2, 414, 400, 3, 2, 2, 2, 414,
	401, 3, 2, 2, 2, 414, 402, 3, 2, 2, 2, 414, 403, 3, 2, 2, 2, 414, 404,
	3, 2, 2, 2, 414, 405, 3, 2, 2, 2, 414, 406, 3, 2, 2, 2, 414, 407, 3, 2,
	2, 2, 414, 408, 3, 2, 2, 2, 414, 409, 3, 2, 2, 2, 414, 410, 3, 2, 2, 2,
	414, 411, 3, 2, 2, 2, 414, 412, 3, 2, 2, 2, 414, 413, 3, 2, 2, 2, 415,
	85, 3, 2, 2, 2, 416, 418, 7, 47, 2, 2, 417, 419, 7, 39, 2, 2, 418, 417,
	3, 2, 2, 2, 418, 419, 3, 2, 2, 2, 419, 420, 3, 2, 2, 2, 420, 421, 7, 48,
	2, 2, 421, 87, 3, 2, 2, 2, 422, 423, 7, 47, 2, 2, 423, 424, 9, 7, 2, 2,
	424, 425, 7, 48, 2, 2, 425, 89, 3, 2, 2, 2, 426, 429, 5, 84, 43, 2, 427,
	429, 7, 88, 2, 2, 428, 426, 3, 2, 2, 2, 428, 427, 3, 2, 2, 2, 429, 432,
	3, 2, 2, 2, 430, 433, 7, 59, 2, 2, 431, 433, 5, 86, 44, 2, 432, 430, 3,
	2, 2, 2, 432, 431, 3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 91, 3, 2, 2,
	2, 434, 435, 7, 88, 2, 2, 435, 436, 5, 90, 46, 2, 436, 93, 3, 2, 2, 2,
	437, 438, 7, 88, 2, 2, 438, 448, 7, 54, 2, 2, 439, 444, 5, 92, 47, 2, 440,
	441, 7, 49, 2, 2, 441, 443, 5, 92, 47, 2, 442, 440, 3, 2, 2, 2, 443, 446,
	3, 2, 2, 2, 444, 442, 3, 2, 2, 2, 444, 445, 3, 2, 2, 2, 445, 448, 3, 2,
	2, 2, 446, 444, 3, 2, 2, 2, 447, 437, 3, 2, 2, 2, 447, 439, 3, 2, 2, 2,
	448, 95, 3, 2, 2, 2, 449, 451, 7, 7, 2, 2, 450, 452, 7, 88, 2, 2, 451,
	450, 3, 2, 2, 2, 451, 452, 3, 2, 2, 2, 452, 453, 3, 2, 2, 2, 453, 455,
	7, 43, 2, 2, 454, 456, 5, 94, 48, 2, 455, 454, 3, 2, 2, 2, 455, 456, 3,
	2, 2, 2, 456, 457, 3, 2, 2, 2, 457, 459, 7, 44, 2, 2, 458, 460, 5, 90,
	46, 2, 459, 458, 3, 2, 2, 2, 459, 460, 3, 2, 2, 2, 460, 461, 3, 2, 2, 2,
	461, 462, 5, 80, 41, 2, 462, 97, 3, 2, 2, 2, 463, 464, 9, 8, 2, 2, 464,
	99, 3, 2, 2, 2, 465, 466, 5, 98, 50, 2, 466, 468, 5, 4, 3, 2, 467, 469,
	5, 90, 46, 2, 468, 467, 3, 2, 2, 2, 468, 469, 3, 2, 2, 2, 469, 471, 3,
	2, 2, 2, 470, 472, 5, 34, 18, 2, 471, 470, 3, 2, 2, 2, 471, 472, 3, 2,
	2, 2, 472, 101, 3, 2, 2, 2, 473, 477, 7, 45, 2, 2, 474, 476, 5, 92, 47,
	2, 475, 474, 3, 2, 2, 2, 476, 479, 3, 2, 2, 2, 477, 475, 3, 2, 2, 2, 477,
	478, 3, 2, 2, 2, 478, 480, 3, 2, 2, 2, 479, 477, 3, 2, 2, 2, 480, 481,
	7, 46, 2, 2, 481, 103, 3, 2, 2, 2, 482, 483, 7, 8, 2, 2, 483, 484, 7, 88,
	2, 2, 484, 485, 5, 102, 52, 2, 485, 105, 3, 2, 2, 2, 51, 109, 118, 125,
	133, 142, 147, 153, 181, 189, 193, 214, 223, 226, 233, 240, 245, 249, 260,
	273, 275, 282, 287, 292, 296, 307, 322, 333, 339, 343, 349, 368, 374, 379,
	389, 392, 396, 398, 414, 418, 428, 432, 444, 447, 451, 455, 459, 468, 471,
	477,
}
var literalNames = []string{
	"", "'let'", "'set'", "'const'", "'delete'", "'func'", "'obj'", "", "",
	"'float'", "'f32'", "'f64'", "'string'", "'bstr'", "'byte'", "'bool'",
	"'symlink'", "'any'", "'true'", "'false'", "'null'", "'self'", "'while'",
	"'loop'", "'times'", "'each'", "'if'", "'else'", "'elseif'", "'switch'",
	"'case'", "'default'", "'continue'", "'return'", "'break'", "'escape'",
	"'typeof'", "", "", "", "", "'('", "')'", "'{'", "'}'", "'['", "']'", "','",
	"'.'", "';'", "'@'", "'::'", "", "'='", "'>'", "'<'", "'!'", "'?'", "':'",
	"'=='", "'<='", "'>='", "'!='", "'&&'", "'||'", "'++'", "'--'", "'+'",
	"'-'", "'*'", "'/'", "'&'", "'|'", "'%'", "'$'", "'<<'", "'>>'", "'..'",
}
var symbolicNames = []string{
	"", "LET", "SET", "CONST", "DELETE", "FUNC", "OBJ", "INT", "UINT", "FLOAT",
	"F32", "F64", "STRING", "BSTR", "BYTE", "BOOL", "SYMLINK", "ANY", "TRUE",
	"FALSE", "NULL", "SELF", "WHILE", "LOOP", "TIMES", "EACH", "IF", "ELSE",
	"ELIF", "SWITCH", "CASE", "DEFAULT", "CONTINUE", "RETURN", "BREAK", "ESCAPE",
	"TYPEOF", "IntegerLiteral", "FloatingPointLiteral", "StringLiteral", "TextBlock",
	"LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "COMMA", "DOT",
	"SEMI", "AT", "DOUBLECOLON", "ELLIPSIS", "ASSIGN", "GT", "LT", "NOT", "QUESTION",
	"COLON", "EQUAL", "LE", "GE", "NOTEQUAL", "AND", "OR", "INC", "DEC", "ADD",
	"SUB", "MUL", "DIV", "BITAND", "BITOR", "MOD", "ADDR", "LSHIFT", "RSHIFT",
	"BITCONCAT", "ADD_ASSIGN", "SUB_ASSIGN", "DIV_ASSIGN", "MUL_ASSIGN", "ADDR_ASSIGN",
	"WSPACE", "COMMENT", "MULTICOMMENT", "Identifier",
}

var ruleNames = []string{
	"start", "identifiers", "annotation", "externalDecl", "op", "arrayMembers",
	"arrayLiteral", "arrayMemberRef", "value", "typeofOperation", "negateOperation",
	"incDecOP", "prefixOperation", "suffixOperation", "operationWithReturn",
	"assignmentOperator", "assignmentOp", "assignment", "returnStatement",
	"callArgs", "funcCall", "comparator", "condition", "elseBlock", "elseifBlock",
	"ifStatement", "namespaceAccess", "memberAccess", "timesDecl", "eachDecl",
	"whileDecl", "loopDecl", "loopStatement", "breakStatement", "switchStatement",
	"cases", "switchBody", "deleteOp", "statement", "body", "function", "builtin",
	"arrayDecl", "arrayAccess", "typeKey", "argument", "arguments", "funcDecl",
	"declarator", "decl", "classBody", "classDecl",
}

type BridgeParser struct {
	*antlr.BaseParser
}

// NewBridgeParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *BridgeParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewBridgeParser(input antlr.TokenStream) *BridgeParser {
	this := new(BridgeParser)
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
	this.GrammarFileName = "Bridge.g4"

	return this
}

// BridgeParser tokens.
const (
	BridgeParserEOF                  = antlr.TokenEOF
	BridgeParserLET                  = 1
	BridgeParserSET                  = 2
	BridgeParserCONST                = 3
	BridgeParserDELETE               = 4
	BridgeParserFUNC                 = 5
	BridgeParserOBJ                  = 6
	BridgeParserINT                  = 7
	BridgeParserUINT                 = 8
	BridgeParserFLOAT                = 9
	BridgeParserF32                  = 10
	BridgeParserF64                  = 11
	BridgeParserSTRING               = 12
	BridgeParserBSTR                 = 13
	BridgeParserBYTE                 = 14
	BridgeParserBOOL                 = 15
	BridgeParserSYMLINK              = 16
	BridgeParserANY                  = 17
	BridgeParserTRUE                 = 18
	BridgeParserFALSE                = 19
	BridgeParserNULL                 = 20
	BridgeParserSELF                 = 21
	BridgeParserWHILE                = 22
	BridgeParserLOOP                 = 23
	BridgeParserTIMES                = 24
	BridgeParserEACH                 = 25
	BridgeParserIF                   = 26
	BridgeParserELSE                 = 27
	BridgeParserELIF                 = 28
	BridgeParserSWITCH               = 29
	BridgeParserCASE                 = 30
	BridgeParserDEFAULT              = 31
	BridgeParserCONTINUE             = 32
	BridgeParserRETURN               = 33
	BridgeParserBREAK                = 34
	BridgeParserESCAPE               = 35
	BridgeParserTYPEOF               = 36
	BridgeParserIntegerLiteral       = 37
	BridgeParserFloatingPointLiteral = 38
	BridgeParserStringLiteral        = 39
	BridgeParserTextBlock            = 40
	BridgeParserLPAREN               = 41
	BridgeParserRPAREN               = 42
	BridgeParserLBRACE               = 43
	BridgeParserRBRACE               = 44
	BridgeParserLBRACK               = 45
	BridgeParserRBRACK               = 46
	BridgeParserCOMMA                = 47
	BridgeParserDOT                  = 48
	BridgeParserSEMI                 = 49
	BridgeParserAT                   = 50
	BridgeParserDOUBLECOLON          = 51
	BridgeParserELLIPSIS             = 52
	BridgeParserASSIGN               = 53
	BridgeParserGT                   = 54
	BridgeParserLT                   = 55
	BridgeParserNOT                  = 56
	BridgeParserQUESTION             = 57
	BridgeParserCOLON                = 58
	BridgeParserEQUAL                = 59
	BridgeParserLE                   = 60
	BridgeParserGE                   = 61
	BridgeParserNOTEQUAL             = 62
	BridgeParserAND                  = 63
	BridgeParserOR                   = 64
	BridgeParserINC                  = 65
	BridgeParserDEC                  = 66
	BridgeParserADD                  = 67
	BridgeParserSUB                  = 68
	BridgeParserMUL                  = 69
	BridgeParserDIV                  = 70
	BridgeParserBITAND               = 71
	BridgeParserBITOR                = 72
	BridgeParserMOD                  = 73
	BridgeParserADDR                 = 74
	BridgeParserLSHIFT               = 75
	BridgeParserRSHIFT               = 76
	BridgeParserBITCONCAT            = 77
	BridgeParserADD_ASSIGN           = 78
	BridgeParserSUB_ASSIGN           = 79
	BridgeParserDIV_ASSIGN           = 80
	BridgeParserMUL_ASSIGN           = 81
	BridgeParserADDR_ASSIGN          = 82
	BridgeParserWSPACE               = 83
	BridgeParserCOMMENT              = 84
	BridgeParserMULTICOMMENT         = 85
	BridgeParserIdentifier           = 86
)

// BridgeParser rules.
const (
	BridgeParserRULE_start               = 0
	BridgeParserRULE_identifiers         = 1
	BridgeParserRULE_annotation          = 2
	BridgeParserRULE_externalDecl        = 3
	BridgeParserRULE_op                  = 4
	BridgeParserRULE_arrayMembers        = 5
	BridgeParserRULE_arrayLiteral        = 6
	BridgeParserRULE_arrayMemberRef      = 7
	BridgeParserRULE_value               = 8
	BridgeParserRULE_typeofOperation     = 9
	BridgeParserRULE_negateOperation     = 10
	BridgeParserRULE_incDecOP            = 11
	BridgeParserRULE_prefixOperation     = 12
	BridgeParserRULE_suffixOperation     = 13
	BridgeParserRULE_operationWithReturn = 14
	BridgeParserRULE_assignmentOperator  = 15
	BridgeParserRULE_assignmentOp        = 16
	BridgeParserRULE_assignment          = 17
	BridgeParserRULE_returnStatement     = 18
	BridgeParserRULE_callArgs            = 19
	BridgeParserRULE_funcCall            = 20
	BridgeParserRULE_comparator          = 21
	BridgeParserRULE_condition           = 22
	BridgeParserRULE_elseBlock           = 23
	BridgeParserRULE_elseifBlock         = 24
	BridgeParserRULE_ifStatement         = 25
	BridgeParserRULE_namespaceAccess     = 26
	BridgeParserRULE_memberAccess        = 27
	BridgeParserRULE_timesDecl           = 28
	BridgeParserRULE_eachDecl            = 29
	BridgeParserRULE_whileDecl           = 30
	BridgeParserRULE_loopDecl            = 31
	BridgeParserRULE_loopStatement       = 32
	BridgeParserRULE_breakStatement      = 33
	BridgeParserRULE_switchStatement     = 34
	BridgeParserRULE_cases               = 35
	BridgeParserRULE_switchBody          = 36
	BridgeParserRULE_deleteOp            = 37
	BridgeParserRULE_statement           = 38
	BridgeParserRULE_body                = 39
	BridgeParserRULE_function            = 40
	BridgeParserRULE_builtin             = 41
	BridgeParserRULE_arrayDecl           = 42
	BridgeParserRULE_arrayAccess         = 43
	BridgeParserRULE_typeKey             = 44
	BridgeParserRULE_argument            = 45
	BridgeParserRULE_arguments           = 46
	BridgeParserRULE_funcDecl            = 47
	BridgeParserRULE_declarator          = 48
	BridgeParserRULE_decl                = 49
	BridgeParserRULE_classBody           = 50
	BridgeParserRULE_classDecl           = 51
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(BridgeParserEOF, 0)
}

func (s *StartContext) AllExternalDecl() []IExternalDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExternalDeclContext)(nil)).Elem())
	var tst = make([]IExternalDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExternalDeclContext)
		}
	}

	return tst
}

func (s *StartContext) ExternalDecl(i int) IExternalDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExternalDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExternalDeclContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BridgeParserRULE_start)
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
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BridgeParserLET)|(1<<BridgeParserSET)|(1<<BridgeParserCONST)|(1<<BridgeParserFUNC)|(1<<BridgeParserOBJ))) != 0) || _la == BridgeParserAT {
		{
			p.SetState(104)
			p.ExternalDecl()
		}

		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(109)
		p.Match(BridgeParserEOF)
	}

	return localctx
}

// IIdentifiersContext is an interface to support dynamic dispatch.
type IIdentifiersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifiersContext differentiates from other interfaces.
	IsIdentifiersContext()
}

type IdentifiersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifiersContext() *IdentifiersContext {
	var p = new(IdentifiersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_identifiers
	return p
}

func (*IdentifiersContext) IsIdentifiersContext() {}

func NewIdentifiersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifiersContext {
	var p = new(IdentifiersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_identifiers

	return p
}

func (s *IdentifiersContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifiersContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(BridgeParserIdentifier)
}

func (s *IdentifiersContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, i)
}

func (s *IdentifiersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BridgeParserCOMMA)
}

func (s *IdentifiersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BridgeParserCOMMA, i)
}

func (s *IdentifiersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifiersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifiersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterIdentifiers(s)
	}
}

func (s *IdentifiersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitIdentifiers(s)
	}
}

func (s *IdentifiersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitIdentifiers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) Identifiers() (localctx IIdentifiersContext) {
	localctx = NewIdentifiersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BridgeParserRULE_identifiers)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(BridgeParserIdentifier)
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(112)
				p.Match(BridgeParserCOMMA)
			}
			{
				p.SetState(113)
				p.Match(BridgeParserIdentifier)
			}

		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IAnnotationContext is an interface to support dynamic dispatch.
type IAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotationContext differentiates from other interfaces.
	IsAnnotationContext()
}

type AnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_annotation
	return p
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationContext) AT() antlr.TerminalNode {
	return s.GetToken(BridgeParserAT, 0)
}

func (s *AnnotationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, 0)
}

func (s *AnnotationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BridgeParserLPAREN, 0)
}

func (s *AnnotationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BridgeParserRPAREN, 0)
}

func (s *AnnotationContext) Identifiers() IIdentifiersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifiersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifiersContext)
}

func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterAnnotation(s)
	}
}

func (s *AnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitAnnotation(s)
	}
}

func (s *AnnotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitAnnotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BridgeParserRULE_annotation)
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
		p.SetState(119)
		p.Match(BridgeParserAT)
	}
	{
		p.SetState(120)
		p.Match(BridgeParserIdentifier)
	}
	{
		p.SetState(121)
		p.Match(BridgeParserLPAREN)
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BridgeParserIdentifier {
		{
			p.SetState(122)
			p.Identifiers()
		}

	}
	{
		p.SetState(125)
		p.Match(BridgeParserRPAREN)
	}

	return localctx
}

// IExternalDeclContext is an interface to support dynamic dispatch.
type IExternalDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExternalDeclContext differentiates from other interfaces.
	IsExternalDeclContext()
}

type ExternalDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExternalDeclContext() *ExternalDeclContext {
	var p = new(ExternalDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_externalDecl
	return p
}

func (*ExternalDeclContext) IsExternalDeclContext() {}

func NewExternalDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternalDeclContext {
	var p = new(ExternalDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_externalDecl

	return p
}

func (s *ExternalDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ExternalDeclContext) Annotation() IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *ExternalDeclContext) FuncDecl() IFuncDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncDeclContext)
}

func (s *ExternalDeclContext) Decl() IDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ExternalDeclContext) ClassDecl() IClassDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassDeclContext)
}

func (s *ExternalDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExternalDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExternalDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterExternalDecl(s)
	}
}

func (s *ExternalDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitExternalDecl(s)
	}
}

func (s *ExternalDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitExternalDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) ExternalDecl() (localctx IExternalDeclContext) {
	localctx = NewExternalDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BridgeParserRULE_externalDecl)

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

	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BridgeParserAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Annotation()
		}

	case BridgeParserFUNC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.FuncDecl()
		}

	case BridgeParserLET, BridgeParserSET, BridgeParserCONST:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(129)
			p.Decl()
		}

	case BridgeParserOBJ:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(130)
			p.ClassDecl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOpContext is an interface to support dynamic dispatch.
type IOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpContext differentiates from other interfaces.
	IsOpContext()
}

type OpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpContext() *OpContext {
	var p = new(OpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_op
	return p
}

func (*OpContext) IsOpContext() {}

func NewOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpContext {
	var p = new(OpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_op

	return p
}

func (s *OpContext) GetParser() antlr.Parser { return s.parser }

func (s *OpContext) ADD() antlr.TerminalNode {
	return s.GetToken(BridgeParserADD, 0)
}

func (s *OpContext) SUB() antlr.TerminalNode {
	return s.GetToken(BridgeParserSUB, 0)
}

func (s *OpContext) DIV() antlr.TerminalNode {
	return s.GetToken(BridgeParserDIV, 0)
}

func (s *OpContext) MOD() antlr.TerminalNode {
	return s.GetToken(BridgeParserMOD, 0)
}

func (s *OpContext) MUL() antlr.TerminalNode {
	return s.GetToken(BridgeParserMUL, 0)
}

func (s *OpContext) BITCONCAT() antlr.TerminalNode {
	return s.GetToken(BridgeParserBITCONCAT, 0)
}

func (s *OpContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(BridgeParserLSHIFT, 0)
}

func (s *OpContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(BridgeParserRSHIFT, 0)
}

func (s *OpContext) OR() antlr.TerminalNode {
	return s.GetToken(BridgeParserOR, 0)
}

func (s *OpContext) AND() antlr.TerminalNode {
	return s.GetToken(BridgeParserAND, 0)
}

func (s *OpContext) BITAND() antlr.TerminalNode {
	return s.GetToken(BridgeParserBITAND, 0)
}

func (s *OpContext) BITOR() antlr.TerminalNode {
	return s.GetToken(BridgeParserBITOR, 0)
}

func (s *OpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterOp(s)
	}
}

func (s *OpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitOp(s)
	}
}

func (s *OpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) Op() (localctx IOpContext) {
	localctx = NewOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BridgeParserRULE_op)
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
		p.SetState(133)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-63)&-(0x1f+1)) == 0 && ((1<<uint((_la-63)))&((1<<(BridgeParserAND-63))|(1<<(BridgeParserOR-63))|(1<<(BridgeParserADD-63))|(1<<(BridgeParserSUB-63))|(1<<(BridgeParserMUL-63))|(1<<(BridgeParserDIV-63))|(1<<(BridgeParserBITAND-63))|(1<<(BridgeParserBITOR-63))|(1<<(BridgeParserMOD-63))|(1<<(BridgeParserLSHIFT-63))|(1<<(BridgeParserRSHIFT-63))|(1<<(BridgeParserBITCONCAT-63)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArrayMembersContext is an interface to support dynamic dispatch.
type IArrayMembersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayMembersContext differentiates from other interfaces.
	IsArrayMembersContext()
}

type ArrayMembersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayMembersContext() *ArrayMembersContext {
	var p = new(ArrayMembersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_arrayMembers
	return p
}

func (*ArrayMembersContext) IsArrayMembersContext() {}

func NewArrayMembersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayMembersContext {
	var p = new(ArrayMembersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_arrayMembers

	return p
}

func (s *ArrayMembersContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayMembersContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ArrayMembersContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArrayMembersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BridgeParserCOMMA)
}

func (s *ArrayMembersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BridgeParserCOMMA, i)
}

func (s *ArrayMembersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayMembersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayMembersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterArrayMembers(s)
	}
}

func (s *ArrayMembersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitArrayMembers(s)
	}
}

func (s *ArrayMembersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitArrayMembers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) ArrayMembers() (localctx IArrayMembersContext) {
	localctx = NewArrayMembersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BridgeParserRULE_arrayMembers)
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
		p.SetState(135)
		p.value(0)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BridgeParserCOMMA {
		{
			p.SetState(136)
			p.Match(BridgeParserCOMMA)
		}
		{
			p.SetState(137)
			p.value(0)
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_arrayLiteral
	return p
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(BridgeParserLBRACK, 0)
}

func (s *ArrayLiteralContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(BridgeParserRBRACK, 0)
}

func (s *ArrayLiteralContext) ArrayMembers() IArrayMembersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayMembersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayMembersContext)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BridgeParserRULE_arrayLiteral)
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
		p.SetState(143)
		p.Match(BridgeParserLBRACK)
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BridgeParserFUNC)|(1<<BridgeParserTRUE)|(1<<BridgeParserFALSE)|(1<<BridgeParserNULL))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(BridgeParserTYPEOF-36))|(1<<(BridgeParserIntegerLiteral-36))|(1<<(BridgeParserFloatingPointLiteral-36))|(1<<(BridgeParserStringLiteral-36))|(1<<(BridgeParserLPAREN-36))|(1<<(BridgeParserLBRACK-36))|(1<<(BridgeParserNOT-36))|(1<<(BridgeParserINC-36))|(1<<(BridgeParserDEC-36)))) != 0) || _la == BridgeParserADDR || _la == BridgeParserIdentifier {
		{
			p.SetState(144)
			p.ArrayMembers()
		}

	}
	{
		p.SetState(147)
		p.Match(BridgeParserRBRACK)
	}

	return localctx
}

// IArrayMemberRefContext is an interface to support dynamic dispatch.
type IArrayMemberRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayMemberRefContext differentiates from other interfaces.
	IsArrayMemberRefContext()
}

type ArrayMemberRefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayMemberRefContext() *ArrayMemberRefContext {
	var p = new(ArrayMemberRefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_arrayMemberRef
	return p
}

func (*ArrayMemberRefContext) IsArrayMemberRefContext() {}

func NewArrayMemberRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayMemberRefContext {
	var p = new(ArrayMemberRefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_arrayMemberRef

	return p
}

func (s *ArrayMemberRefContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayMemberRefContext) ArrayAccess() IArrayAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayAccessContext)
}

func (s *ArrayMemberRefContext) MemberAccess() IMemberAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberAccessContext)
}

func (s *ArrayMemberRefContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, 0)
}

func (s *ArrayMemberRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayMemberRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayMemberRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterArrayMemberRef(s)
	}
}

func (s *ArrayMemberRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitArrayMemberRef(s)
	}
}

func (s *ArrayMemberRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitArrayMemberRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) ArrayMemberRef() (localctx IArrayMemberRefContext) {
	localctx = NewArrayMemberRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BridgeParserRULE_arrayMemberRef)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(149)
			p.MemberAccess()
		}

	case 2:
		{
			p.SetState(150)
			p.Match(BridgeParserIdentifier)
		}

	}
	{
		p.SetState(153)
		p.ArrayAccess()
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BridgeParserLPAREN, 0)
}

func (s *ValueContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ValueContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BridgeParserRPAREN, 0)
}

func (s *ValueContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ValueContext) OperationWithReturn() IOperationWithReturnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationWithReturnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationWithReturnContext)
}

func (s *ValueContext) FuncCall() IFuncCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *ValueContext) FuncDecl() IFuncDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncDeclContext)
}

func (s *ValueContext) ArrayMemberRef() IArrayMemberRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayMemberRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayMemberRefContext)
}

func (s *ValueContext) MemberAccess() IMemberAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberAccessContext)
}

func (s *ValueContext) ArrayLiteral() IArrayLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *ValueContext) ADDR() antlr.TerminalNode {
	return s.GetToken(BridgeParserADDR, 0)
}

func (s *ValueContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, 0)
}

func (s *ValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(BridgeParserTRUE, 0)
}

func (s *ValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(BridgeParserFALSE, 0)
}

func (s *ValueContext) NULL() antlr.TerminalNode {
	return s.GetToken(BridgeParserNULL, 0)
}

func (s *ValueContext) FloatingPointLiteral() antlr.TerminalNode {
	return s.GetToken(BridgeParserFloatingPointLiteral, 0)
}

func (s *ValueContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(BridgeParserIntegerLiteral, 0)
}

func (s *ValueContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(BridgeParserStringLiteral, 0)
}

func (s *ValueContext) AllOp() []IOpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOpContext)(nil)).Elem())
	var tst = make([]IOpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOpContext)
		}
	}

	return tst
}

func (s *ValueContext) Op(i int) IOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOpContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) Value() (localctx IValueContext) {
	return p.value(0)
}

func (p *BridgeParser) value(_p int) (localctx IValueContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewValueContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IValueContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, BridgeParserRULE_value, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(156)
			p.Match(BridgeParserLPAREN)
		}
		{
			p.SetState(157)
			p.value(0)
		}
		{
			p.SetState(158)
			p.Match(BridgeParserRPAREN)
		}

	case 2:
		{
			p.SetState(160)
			p.Match(BridgeParserLPAREN)
		}
		{
			p.SetState(161)
			p.Condition()
		}
		{
			p.SetState(162)
			p.Match(BridgeParserRPAREN)
		}

	case 3:
		{
			p.SetState(164)
			p.OperationWithReturn()
		}

	case 4:
		{
			p.SetState(165)
			p.FuncCall()
		}

	case 5:
		{
			p.SetState(166)
			p.FuncDecl()
		}

	case 6:
		{
			p.SetState(167)
			p.ArrayMemberRef()
		}

	case 7:
		{
			p.SetState(168)
			p.MemberAccess()
		}

	case 8:
		{
			p.SetState(169)
			p.ArrayLiteral()
		}

	case 9:
		{
			p.SetState(170)
			p.Match(BridgeParserADDR)
		}
		{
			p.SetState(171)
			p.Match(BridgeParserIdentifier)
		}

	case 10:
		{
			p.SetState(172)
			p.Match(BridgeParserTRUE)
		}

	case 11:
		{
			p.SetState(173)
			p.Match(BridgeParserFALSE)
		}

	case 12:
		{
			p.SetState(174)
			p.Match(BridgeParserNULL)
		}

	case 13:
		{
			p.SetState(175)
			p.Match(BridgeParserFloatingPointLiteral)
		}

	case 14:
		{
			p.SetState(176)
			p.Match(BridgeParserIntegerLiteral)
		}

	case 15:
		{
			p.SetState(177)
			p.Match(BridgeParserStringLiteral)
		}

	case 16:
		{
			p.SetState(178)
			p.Match(BridgeParserIdentifier)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewValueContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, BridgeParserRULE_value)
			p.SetState(181)

			if !(p.Precpred(p.GetParserRuleContext(), 14)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
			}
			p.SetState(185)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(182)
						p.Op()
					}
					{
						p.SetState(183)
						p.value(0)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(187)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
			}

		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// ITypeofOperationContext is an interface to support dynamic dispatch.
type ITypeofOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeofOperationContext differentiates from other interfaces.
	IsTypeofOperationContext()
}

type TypeofOperationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeofOperationContext() *TypeofOperationContext {
	var p = new(TypeofOperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_typeofOperation
	return p
}

func (*TypeofOperationContext) IsTypeofOperationContext() {}

func NewTypeofOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeofOperationContext {
	var p = new(TypeofOperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_typeofOperation

	return p
}

func (s *TypeofOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeofOperationContext) TYPEOF() antlr.TerminalNode {
	return s.GetToken(BridgeParserTYPEOF, 0)
}

func (s *TypeofOperationContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *TypeofOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeofOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeofOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterTypeofOperation(s)
	}
}

func (s *TypeofOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitTypeofOperation(s)
	}
}

func (s *TypeofOperationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitTypeofOperation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) TypeofOperation() (localctx ITypeofOperationContext) {
	localctx = NewTypeofOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BridgeParserRULE_typeofOperation)

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
		p.SetState(194)
		p.Match(BridgeParserTYPEOF)
	}
	{
		p.SetState(195)
		p.value(0)
	}

	return localctx
}

// INegateOperationContext is an interface to support dynamic dispatch.
type INegateOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNegateOperationContext differentiates from other interfaces.
	IsNegateOperationContext()
}

type NegateOperationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNegateOperationContext() *NegateOperationContext {
	var p = new(NegateOperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_negateOperation
	return p
}

func (*NegateOperationContext) IsNegateOperationContext() {}

func NewNegateOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NegateOperationContext {
	var p = new(NegateOperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_negateOperation

	return p
}

func (s *NegateOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *NegateOperationContext) NOT() antlr.TerminalNode {
	return s.GetToken(BridgeParserNOT, 0)
}

func (s *NegateOperationContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *NegateOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegateOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NegateOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterNegateOperation(s)
	}
}

func (s *NegateOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitNegateOperation(s)
	}
}

func (s *NegateOperationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitNegateOperation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) NegateOperation() (localctx INegateOperationContext) {
	localctx = NewNegateOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BridgeParserRULE_negateOperation)

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
		p.SetState(197)
		p.Match(BridgeParserNOT)
	}
	{
		p.SetState(198)
		p.value(0)
	}

	return localctx
}

// IIncDecOPContext is an interface to support dynamic dispatch.
type IIncDecOPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncDecOPContext differentiates from other interfaces.
	IsIncDecOPContext()
}

type IncDecOPContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncDecOPContext() *IncDecOPContext {
	var p = new(IncDecOPContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_incDecOP
	return p
}

func (*IncDecOPContext) IsIncDecOPContext() {}

func NewIncDecOPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncDecOPContext {
	var p = new(IncDecOPContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_incDecOP

	return p
}

func (s *IncDecOPContext) GetParser() antlr.Parser { return s.parser }

func (s *IncDecOPContext) INC() antlr.TerminalNode {
	return s.GetToken(BridgeParserINC, 0)
}

func (s *IncDecOPContext) DEC() antlr.TerminalNode {
	return s.GetToken(BridgeParserDEC, 0)
}

func (s *IncDecOPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncDecOPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncDecOPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterIncDecOP(s)
	}
}

func (s *IncDecOPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitIncDecOP(s)
	}
}

func (s *IncDecOPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitIncDecOP(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) IncDecOP() (localctx IIncDecOPContext) {
	localctx = NewIncDecOPContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BridgeParserRULE_incDecOP)
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
		p.SetState(200)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BridgeParserINC || _la == BridgeParserDEC) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPrefixOperationContext is an interface to support dynamic dispatch.
type IPrefixOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefixOperationContext differentiates from other interfaces.
	IsPrefixOperationContext()
}

type PrefixOperationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixOperationContext() *PrefixOperationContext {
	var p = new(PrefixOperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_prefixOperation
	return p
}

func (*PrefixOperationContext) IsPrefixOperationContext() {}

func NewPrefixOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixOperationContext {
	var p = new(PrefixOperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_prefixOperation

	return p
}

func (s *PrefixOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixOperationContext) IncDecOP() IIncDecOPContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncDecOPContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIncDecOPContext)
}

func (s *PrefixOperationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, 0)
}

func (s *PrefixOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefixOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterPrefixOperation(s)
	}
}

func (s *PrefixOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitPrefixOperation(s)
	}
}

func (s *PrefixOperationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitPrefixOperation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) PrefixOperation() (localctx IPrefixOperationContext) {
	localctx = NewPrefixOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BridgeParserRULE_prefixOperation)

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
		p.SetState(202)
		p.IncDecOP()
	}
	{
		p.SetState(203)
		p.Match(BridgeParserIdentifier)
	}

	return localctx
}

// ISuffixOperationContext is an interface to support dynamic dispatch.
type ISuffixOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSuffixOperationContext differentiates from other interfaces.
	IsSuffixOperationContext()
}

type SuffixOperationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySuffixOperationContext() *SuffixOperationContext {
	var p = new(SuffixOperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_suffixOperation
	return p
}

func (*SuffixOperationContext) IsSuffixOperationContext() {}

func NewSuffixOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SuffixOperationContext {
	var p = new(SuffixOperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_suffixOperation

	return p
}

func (s *SuffixOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *SuffixOperationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, 0)
}

func (s *SuffixOperationContext) IncDecOP() IIncDecOPContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncDecOPContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIncDecOPContext)
}

func (s *SuffixOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuffixOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SuffixOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterSuffixOperation(s)
	}
}

func (s *SuffixOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitSuffixOperation(s)
	}
}

func (s *SuffixOperationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitSuffixOperation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) SuffixOperation() (localctx ISuffixOperationContext) {
	localctx = NewSuffixOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BridgeParserRULE_suffixOperation)

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
		p.SetState(205)
		p.Match(BridgeParserIdentifier)
	}
	{
		p.SetState(206)
		p.IncDecOP()
	}

	return localctx
}

// IOperationWithReturnContext is an interface to support dynamic dispatch.
type IOperationWithReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationWithReturnContext differentiates from other interfaces.
	IsOperationWithReturnContext()
}

type OperationWithReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationWithReturnContext() *OperationWithReturnContext {
	var p = new(OperationWithReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_operationWithReturn
	return p
}

func (*OperationWithReturnContext) IsOperationWithReturnContext() {}

func NewOperationWithReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationWithReturnContext {
	var p = new(OperationWithReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_operationWithReturn

	return p
}

func (s *OperationWithReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationWithReturnContext) TypeofOperation() ITypeofOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeofOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeofOperationContext)
}

func (s *OperationWithReturnContext) NegateOperation() INegateOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INegateOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INegateOperationContext)
}

func (s *OperationWithReturnContext) PrefixOperation() IPrefixOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixOperationContext)
}

func (s *OperationWithReturnContext) SuffixOperation() ISuffixOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISuffixOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISuffixOperationContext)
}

func (s *OperationWithReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationWithReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationWithReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterOperationWithReturn(s)
	}
}

func (s *OperationWithReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitOperationWithReturn(s)
	}
}

func (s *OperationWithReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitOperationWithReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) OperationWithReturn() (localctx IOperationWithReturnContext) {
	localctx = NewOperationWithReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BridgeParserRULE_operationWithReturn)

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

	p.SetState(212)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BridgeParserTYPEOF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(208)
			p.TypeofOperation()
		}

	case BridgeParserNOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(209)
			p.NegateOperation()
		}

	case BridgeParserINC, BridgeParserDEC:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(210)
			p.PrefixOperation()
		}

	case BridgeParserIdentifier:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(211)
			p.SuffixOperation()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssignmentOperatorContext is an interface to support dynamic dispatch.
type IAssignmentOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentOperatorContext differentiates from other interfaces.
	IsAssignmentOperatorContext()
}

type AssignmentOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentOperatorContext() *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_assignmentOperator
	return p
}

func (*AssignmentOperatorContext) IsAssignmentOperatorContext() {}

func NewAssignmentOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_assignmentOperator

	return p
}

func (s *AssignmentOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentOperatorContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(BridgeParserASSIGN, 0)
}

func (s *AssignmentOperatorContext) ADD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(BridgeParserADD_ASSIGN, 0)
}

func (s *AssignmentOperatorContext) SUB_ASSIGN() antlr.TerminalNode {
	return s.GetToken(BridgeParserSUB_ASSIGN, 0)
}

func (s *AssignmentOperatorContext) MUL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(BridgeParserMUL_ASSIGN, 0)
}

func (s *AssignmentOperatorContext) DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(BridgeParserDIV_ASSIGN, 0)
}

func (s *AssignmentOperatorContext) ADDR_ASSIGN() antlr.TerminalNode {
	return s.GetToken(BridgeParserADDR_ASSIGN, 0)
}

func (s *AssignmentOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterAssignmentOperator(s)
	}
}

func (s *AssignmentOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitAssignmentOperator(s)
	}
}

func (s *AssignmentOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitAssignmentOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) AssignmentOperator() (localctx IAssignmentOperatorContext) {
	localctx = NewAssignmentOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BridgeParserRULE_assignmentOperator)
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
		p.SetState(214)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(BridgeParserASSIGN-53))|(1<<(BridgeParserADD_ASSIGN-53))|(1<<(BridgeParserSUB_ASSIGN-53))|(1<<(BridgeParserDIV_ASSIGN-53))|(1<<(BridgeParserMUL_ASSIGN-53))|(1<<(BridgeParserADDR_ASSIGN-53)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAssignmentOpContext is an interface to support dynamic dispatch.
type IAssignmentOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentOpContext differentiates from other interfaces.
	IsAssignmentOpContext()
}

type AssignmentOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentOpContext() *AssignmentOpContext {
	var p = new(AssignmentOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_assignmentOp
	return p
}

func (*AssignmentOpContext) IsAssignmentOpContext() {}

func NewAssignmentOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentOpContext {
	var p = new(AssignmentOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_assignmentOp

	return p
}

func (s *AssignmentOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentOpContext) AssignmentOperator() IAssignmentOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentOperatorContext)
}

func (s *AssignmentOpContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *AssignmentOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterAssignmentOp(s)
	}
}

func (s *AssignmentOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitAssignmentOp(s)
	}
}

func (s *AssignmentOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitAssignmentOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) AssignmentOp() (localctx IAssignmentOpContext) {
	localctx = NewAssignmentOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BridgeParserRULE_assignmentOp)

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
		p.SetState(216)
		p.AssignmentOperator()
	}
	{
		p.SetState(217)
		p.value(0)
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) AssignmentOp() IAssignmentOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentOpContext)
}

func (s *AssignmentContext) MemberAccess() IMemberAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberAccessContext)
}

func (s *AssignmentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, 0)
}

func (s *AssignmentContext) ArrayAccess() IArrayAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayAccessContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BridgeParserRULE_assignment)
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
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(219)
			p.MemberAccess()
		}

	case 2:
		{
			p.SetState(220)
			p.Match(BridgeParserIdentifier)
		}

	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BridgeParserLBRACK {
		{
			p.SetState(223)
			p.ArrayAccess()
		}

	}

	{
		p.SetState(226)
		p.AssignmentOp()
	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(BridgeParserRETURN, 0)
}

func (s *ReturnStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ReturnStatementContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BridgeParserRULE_returnStatement)

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
		p.SetState(228)
		p.Match(BridgeParserRETURN)
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(229)
			p.Statement()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(230)
			p.value(0)
		}

	}

	return localctx
}

// ICallArgsContext is an interface to support dynamic dispatch.
type ICallArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallArgsContext differentiates from other interfaces.
	IsCallArgsContext()
}

type CallArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallArgsContext() *CallArgsContext {
	var p = new(CallArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_callArgs
	return p
}

func (*CallArgsContext) IsCallArgsContext() {}

func NewCallArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallArgsContext {
	var p = new(CallArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_callArgs

	return p
}

func (s *CallArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *CallArgsContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *CallArgsContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *CallArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BridgeParserCOMMA)
}

func (s *CallArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BridgeParserCOMMA, i)
}

func (s *CallArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterCallArgs(s)
	}
}

func (s *CallArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitCallArgs(s)
	}
}

func (s *CallArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitCallArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) CallArgs() (localctx ICallArgsContext) {
	localctx = NewCallArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BridgeParserRULE_callArgs)
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
		p.SetState(233)
		p.value(0)
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BridgeParserCOMMA {
		{
			p.SetState(234)
			p.Match(BridgeParserCOMMA)
		}
		{
			p.SetState(235)
			p.value(0)
		}

		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFuncCallContext is an interface to support dynamic dispatch.
type IFuncCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncCallContext differentiates from other interfaces.
	IsFuncCallContext()
}

type FuncCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallContext() *FuncCallContext {
	var p = new(FuncCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_funcCall
	return p
}

func (*FuncCallContext) IsFuncCallContext() {}

func NewFuncCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallContext {
	var p = new(FuncCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_funcCall

	return p
}

func (s *FuncCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BridgeParserLPAREN, 0)
}

func (s *FuncCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BridgeParserRPAREN, 0)
}

func (s *FuncCallContext) NamespaceAccess() INamespaceAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespaceAccessContext)
}

func (s *FuncCallContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, 0)
}

func (s *FuncCallContext) CallArgs() ICallArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallArgsContext)
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

func (s *FuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) FuncCall() (localctx IFuncCallContext) {
	localctx = NewFuncCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BridgeParserRULE_funcCall)
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
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(241)
			p.NamespaceAccess()
		}

	case 2:
		{
			p.SetState(242)
			p.Match(BridgeParserIdentifier)
		}

	}
	{
		p.SetState(245)
		p.Match(BridgeParserLPAREN)
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BridgeParserFUNC)|(1<<BridgeParserTRUE)|(1<<BridgeParserFALSE)|(1<<BridgeParserNULL))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(BridgeParserTYPEOF-36))|(1<<(BridgeParserIntegerLiteral-36))|(1<<(BridgeParserFloatingPointLiteral-36))|(1<<(BridgeParserStringLiteral-36))|(1<<(BridgeParserLPAREN-36))|(1<<(BridgeParserLBRACK-36))|(1<<(BridgeParserNOT-36))|(1<<(BridgeParserINC-36))|(1<<(BridgeParserDEC-36)))) != 0) || _la == BridgeParserADDR || _la == BridgeParserIdentifier {
		{
			p.SetState(246)
			p.CallArgs()
		}

	}
	{
		p.SetState(249)
		p.Match(BridgeParserRPAREN)
	}

	return localctx
}

// IComparatorContext is an interface to support dynamic dispatch.
type IComparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparatorContext differentiates from other interfaces.
	IsComparatorContext()
}

type ComparatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparatorContext() *ComparatorContext {
	var p = new(ComparatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_comparator
	return p
}

func (*ComparatorContext) IsComparatorContext() {}

func NewComparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorContext {
	var p = new(ComparatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_comparator

	return p
}

func (s *ComparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparatorContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(BridgeParserEQUAL, 0)
}

func (s *ComparatorContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(BridgeParserNOTEQUAL, 0)
}

func (s *ComparatorContext) GT() antlr.TerminalNode {
	return s.GetToken(BridgeParserGT, 0)
}

func (s *ComparatorContext) GE() antlr.TerminalNode {
	return s.GetToken(BridgeParserGE, 0)
}

func (s *ComparatorContext) LT() antlr.TerminalNode {
	return s.GetToken(BridgeParserLT, 0)
}

func (s *ComparatorContext) LE() antlr.TerminalNode {
	return s.GetToken(BridgeParserLE, 0)
}

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterComparator(s)
	}
}

func (s *ComparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitComparator(s)
	}
}

func (s *ComparatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitComparator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) Comparator() (localctx IComparatorContext) {
	localctx = NewComparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BridgeParserRULE_comparator)
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
		p.SetState(251)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(BridgeParserGT-54))|(1<<(BridgeParserLT-54))|(1<<(BridgeParserEQUAL-54))|(1<<(BridgeParserLE-54))|(1<<(BridgeParserGE-54))|(1<<(BridgeParserNOTEQUAL-54)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ConditionContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ConditionContext) Comparator() IComparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparatorContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BridgeParserRULE_condition)

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

	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.value(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)
			p.value(0)
		}
		{
			p.SetState(255)
			p.Comparator()
		}
		{
			p.SetState(256)
			p.value(0)
		}

	}

	return localctx
}

// IElseBlockContext is an interface to support dynamic dispatch.
type IElseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseBlockContext differentiates from other interfaces.
	IsElseBlockContext()
}

type ElseBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseBlockContext() *ElseBlockContext {
	var p = new(ElseBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_elseBlock
	return p
}

func (*ElseBlockContext) IsElseBlockContext() {}

func NewElseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseBlockContext {
	var p = new(ElseBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_elseBlock

	return p
}

func (s *ElseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseBlockContext) ELSE() antlr.TerminalNode {
	return s.GetToken(BridgeParserELSE, 0)
}

func (s *ElseBlockContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *ElseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterElseBlock(s)
	}
}

func (s *ElseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitElseBlock(s)
	}
}

func (s *ElseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitElseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) ElseBlock() (localctx IElseBlockContext) {
	localctx = NewElseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BridgeParserRULE_elseBlock)

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
		p.SetState(260)
		p.Match(BridgeParserELSE)
	}
	{
		p.SetState(261)
		p.Body()
	}

	return localctx
}

// IElseifBlockContext is an interface to support dynamic dispatch.
type IElseifBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseifBlockContext differentiates from other interfaces.
	IsElseifBlockContext()
}

type ElseifBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseifBlockContext() *ElseifBlockContext {
	var p = new(ElseifBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_elseifBlock
	return p
}

func (*ElseifBlockContext) IsElseifBlockContext() {}

func NewElseifBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseifBlockContext {
	var p = new(ElseifBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_elseifBlock

	return p
}

func (s *ElseifBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseifBlockContext) ELSE() antlr.TerminalNode {
	return s.GetToken(BridgeParserELSE, 0)
}

func (s *ElseifBlockContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *ElseifBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseifBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseifBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterElseifBlock(s)
	}
}

func (s *ElseifBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitElseifBlock(s)
	}
}

func (s *ElseifBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitElseifBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) ElseifBlock() (localctx IElseifBlockContext) {
	localctx = NewElseifBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, BridgeParserRULE_elseifBlock)

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
		p.SetState(263)
		p.Match(BridgeParserELSE)
	}
	{
		p.SetState(264)
		p.IfStatement()
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(BridgeParserIF, 0)
}

func (s *IfStatementContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *IfStatementContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *IfStatementContext) AllElseifBlock() []IElseifBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElseifBlockContext)(nil)).Elem())
	var tst = make([]IElseifBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElseifBlockContext)
		}
	}

	return tst
}

func (s *IfStatementContext) ElseifBlock(i int) IElseifBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseifBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElseifBlockContext)
}

func (s *IfStatementContext) AllElseBlock() []IElseBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElseBlockContext)(nil)).Elem())
	var tst = make([]IElseBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElseBlockContext)
		}
	}

	return tst
}

func (s *IfStatementContext) ElseBlock(i int) IElseBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElseBlockContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, BridgeParserRULE_ifStatement)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(BridgeParserIF)
	}
	{
		p.SetState(267)
		p.Condition()
	}
	{
		p.SetState(268)
		p.Body()
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(271)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(269)
					p.ElseifBlock()
				}

			case 2:
				{
					p.SetState(270)
					p.ElseBlock()
				}

			}

		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

// INamespaceAccessContext is an interface to support dynamic dispatch.
type INamespaceAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespaceAccessContext differentiates from other interfaces.
	IsNamespaceAccessContext()
}

type NamespaceAccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceAccessContext() *NamespaceAccessContext {
	var p = new(NamespaceAccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_namespaceAccess
	return p
}

func (*NamespaceAccessContext) IsNamespaceAccessContext() {}

func NewNamespaceAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceAccessContext {
	var p = new(NamespaceAccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_namespaceAccess

	return p
}

func (s *NamespaceAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceAccessContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(BridgeParserIdentifier)
}

func (s *NamespaceAccessContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, i)
}

func (s *NamespaceAccessContext) DOUBLECOLON() antlr.TerminalNode {
	return s.GetToken(BridgeParserDOUBLECOLON, 0)
}

func (s *NamespaceAccessContext) FuncCall() IFuncCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *NamespaceAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterNamespaceAccess(s)
	}
}

func (s *NamespaceAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitNamespaceAccess(s)
	}
}

func (s *NamespaceAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitNamespaceAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) NamespaceAccess() (localctx INamespaceAccessContext) {
	localctx = NewNamespaceAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BridgeParserRULE_namespaceAccess)

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
		p.SetState(276)
		p.Match(BridgeParserIdentifier)
	}
	{
		p.SetState(277)
		p.Match(BridgeParserDOUBLECOLON)
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(278)
			p.FuncCall()
		}

	case 2:
		{
			p.SetState(279)
			p.Match(BridgeParserIdentifier)
		}

	}

	return localctx
}

// IMemberAccessContext is an interface to support dynamic dispatch.
type IMemberAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberAccessContext differentiates from other interfaces.
	IsMemberAccessContext()
}

type MemberAccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberAccessContext() *MemberAccessContext {
	var p = new(MemberAccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_memberAccess
	return p
}

func (*MemberAccessContext) IsMemberAccessContext() {}

func NewMemberAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberAccessContext {
	var p = new(MemberAccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_memberAccess

	return p
}

func (s *MemberAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberAccessContext) NamespaceAccess() INamespaceAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespaceAccessContext)
}

func (s *MemberAccessContext) AllFuncCall() []IFuncCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncCallContext)(nil)).Elem())
	var tst = make([]IFuncCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncCallContext)
		}
	}

	return tst
}

func (s *MemberAccessContext) FuncCall(i int) IFuncCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *MemberAccessContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(BridgeParserIdentifier)
}

func (s *MemberAccessContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, i)
}

func (s *MemberAccessContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(BridgeParserDOT)
}

func (s *MemberAccessContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(BridgeParserDOT, i)
}

func (s *MemberAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterMemberAccess(s)
	}
}

func (s *MemberAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitMemberAccess(s)
	}
}

func (s *MemberAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitMemberAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) MemberAccess() (localctx IMemberAccessContext) {
	localctx = NewMemberAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BridgeParserRULE_memberAccess)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(282)
			p.NamespaceAccess()
		}

	case 2:
		{
			p.SetState(283)
			p.FuncCall()
		}

	case 3:
		{
			p.SetState(284)
			p.Match(BridgeParserIdentifier)
		}

	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(287)
				p.Match(BridgeParserDOT)
			}
			p.SetState(290)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(288)
					p.FuncCall()
				}

			case 2:
				{
					p.SetState(289)
					p.Match(BridgeParserIdentifier)
				}

			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// ITimesDeclContext is an interface to support dynamic dispatch.
type ITimesDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimesDeclContext differentiates from other interfaces.
	IsTimesDeclContext()
}

type TimesDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimesDeclContext() *TimesDeclContext {
	var p = new(TimesDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_timesDecl
	return p
}

func (*TimesDeclContext) IsTimesDeclContext() {}

func NewTimesDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimesDeclContext {
	var p = new(TimesDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_timesDecl

	return p
}

func (s *TimesDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TimesDeclContext) TIMES() antlr.TerminalNode {
	return s.GetToken(BridgeParserTIMES, 0)
}

func (s *TimesDeclContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *TimesDeclContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *TimesDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimesDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimesDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterTimesDecl(s)
	}
}

func (s *TimesDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitTimesDecl(s)
	}
}

func (s *TimesDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitTimesDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) TimesDecl() (localctx ITimesDeclContext) {
	localctx = NewTimesDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BridgeParserRULE_timesDecl)

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
		p.Match(BridgeParserTIMES)
	}
	{
		p.SetState(297)
		p.value(0)
	}
	{
		p.SetState(298)
		p.Body()
	}

	return localctx
}

// IEachDeclContext is an interface to support dynamic dispatch.
type IEachDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEachDeclContext differentiates from other interfaces.
	IsEachDeclContext()
}

type EachDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEachDeclContext() *EachDeclContext {
	var p = new(EachDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_eachDecl
	return p
}

func (*EachDeclContext) IsEachDeclContext() {}

func NewEachDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EachDeclContext {
	var p = new(EachDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_eachDecl

	return p
}

func (s *EachDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *EachDeclContext) EACH() antlr.TerminalNode {
	return s.GetToken(BridgeParserEACH, 0)
}

func (s *EachDeclContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(BridgeParserIdentifier)
}

func (s *EachDeclContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, i)
}

func (s *EachDeclContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *EachDeclContext) MemberAccess() IMemberAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberAccessContext)
}

func (s *EachDeclContext) ArrayLiteral() IArrayLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *EachDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EachDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EachDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterEachDecl(s)
	}
}

func (s *EachDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitEachDecl(s)
	}
}

func (s *EachDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitEachDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) EachDecl() (localctx IEachDeclContext) {
	localctx = NewEachDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BridgeParserRULE_eachDecl)

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
		p.SetState(300)
		p.Match(BridgeParserEACH)
	}
	{
		p.SetState(301)
		p.Match(BridgeParserIdentifier)
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(302)
			p.MemberAccess()
		}

	case 2:
		{
			p.SetState(303)
			p.Match(BridgeParserIdentifier)
		}

	case 3:
		{
			p.SetState(304)
			p.ArrayLiteral()
		}

	}
	{
		p.SetState(307)
		p.Body()
	}

	return localctx
}

// IWhileDeclContext is an interface to support dynamic dispatch.
type IWhileDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileDeclContext differentiates from other interfaces.
	IsWhileDeclContext()
}

type WhileDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileDeclContext() *WhileDeclContext {
	var p = new(WhileDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_whileDecl
	return p
}

func (*WhileDeclContext) IsWhileDeclContext() {}

func NewWhileDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileDeclContext {
	var p = new(WhileDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_whileDecl

	return p
}

func (s *WhileDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileDeclContext) WHILE() antlr.TerminalNode {
	return s.GetToken(BridgeParserWHILE, 0)
}

func (s *WhileDeclContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *WhileDeclContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *WhileDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterWhileDecl(s)
	}
}

func (s *WhileDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitWhileDecl(s)
	}
}

func (s *WhileDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitWhileDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) WhileDecl() (localctx IWhileDeclContext) {
	localctx = NewWhileDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, BridgeParserRULE_whileDecl)

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
		p.SetState(309)
		p.Match(BridgeParserWHILE)
	}
	{
		p.SetState(310)
		p.Condition()
	}
	{
		p.SetState(311)
		p.Body()
	}

	return localctx
}

// ILoopDeclContext is an interface to support dynamic dispatch.
type ILoopDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopDeclContext differentiates from other interfaces.
	IsLoopDeclContext()
}

type LoopDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopDeclContext() *LoopDeclContext {
	var p = new(LoopDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_loopDecl
	return p
}

func (*LoopDeclContext) IsLoopDeclContext() {}

func NewLoopDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopDeclContext {
	var p = new(LoopDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_loopDecl

	return p
}

func (s *LoopDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopDeclContext) LOOP() antlr.TerminalNode {
	return s.GetToken(BridgeParserLOOP, 0)
}

func (s *LoopDeclContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *LoopDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterLoopDecl(s)
	}
}

func (s *LoopDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitLoopDecl(s)
	}
}

func (s *LoopDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitLoopDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) LoopDecl() (localctx ILoopDeclContext) {
	localctx = NewLoopDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, BridgeParserRULE_loopDecl)

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
		p.SetState(313)
		p.Match(BridgeParserLOOP)
	}
	{
		p.SetState(314)
		p.Body()
	}

	return localctx
}

// ILoopStatementContext is an interface to support dynamic dispatch.
type ILoopStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopStatementContext differentiates from other interfaces.
	IsLoopStatementContext()
}

type LoopStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopStatementContext() *LoopStatementContext {
	var p = new(LoopStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_loopStatement
	return p
}

func (*LoopStatementContext) IsLoopStatementContext() {}

func NewLoopStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopStatementContext {
	var p = new(LoopStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_loopStatement

	return p
}

func (s *LoopStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopStatementContext) LoopDecl() ILoopDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopDeclContext)
}

func (s *LoopStatementContext) WhileDecl() IWhileDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileDeclContext)
}

func (s *LoopStatementContext) EachDecl() IEachDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEachDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEachDeclContext)
}

func (s *LoopStatementContext) TimesDecl() ITimesDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimesDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimesDeclContext)
}

func (s *LoopStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterLoopStatement(s)
	}
}

func (s *LoopStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitLoopStatement(s)
	}
}

func (s *LoopStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitLoopStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) LoopStatement() (localctx ILoopStatementContext) {
	localctx = NewLoopStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, BridgeParserRULE_loopStatement)

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

	p.SetState(320)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BridgeParserLOOP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(316)
			p.LoopDecl()
		}

	case BridgeParserWHILE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(317)
			p.WhileDecl()
		}

	case BridgeParserEACH:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(318)
			p.EachDecl()
		}

	case BridgeParserTIMES:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(319)
			p.TimesDecl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBreakStatementContext is an interface to support dynamic dispatch.
type IBreakStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreakStatementContext differentiates from other interfaces.
	IsBreakStatementContext()
}

type BreakStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatementContext() *BreakStatementContext {
	var p = new(BreakStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_breakStatement
	return p
}

func (*BreakStatementContext) IsBreakStatementContext() {}

func NewBreakStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatementContext {
	var p = new(BreakStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_breakStatement

	return p
}

func (s *BreakStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStatementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(BridgeParserBREAK, 0)
}

func (s *BreakStatementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(BridgeParserCONTINUE, 0)
}

func (s *BreakStatementContext) ESCAPE() antlr.TerminalNode {
	return s.GetToken(BridgeParserESCAPE, 0)
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) BreakStatement() (localctx IBreakStatementContext) {
	localctx = NewBreakStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, BridgeParserRULE_breakStatement)
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
		p.SetState(322)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BridgeParserCONTINUE-32))|(1<<(BridgeParserBREAK-32))|(1<<(BridgeParserESCAPE-32)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISwitchStatementContext is an interface to support dynamic dispatch.
type ISwitchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchStatementContext differentiates from other interfaces.
	IsSwitchStatementContext()
}

type SwitchStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStatementContext() *SwitchStatementContext {
	var p = new(SwitchStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_switchStatement
	return p
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(BridgeParserSWITCH, 0)
}

func (s *SwitchStatementContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *SwitchStatementContext) SwitchBody() ISwitchBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwitchBodyContext)
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitSwitchStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) SwitchStatement() (localctx ISwitchStatementContext) {
	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, BridgeParserRULE_switchStatement)

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
		p.SetState(324)
		p.Match(BridgeParserSWITCH)
	}
	{
		p.SetState(325)
		p.value(0)
	}
	{
		p.SetState(326)
		p.SwitchBody()
	}

	return localctx
}

// ICasesContext is an interface to support dynamic dispatch.
type ICasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCasesContext differentiates from other interfaces.
	IsCasesContext()
}

type CasesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCasesContext() *CasesContext {
	var p = new(CasesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_cases
	return p
}

func (*CasesContext) IsCasesContext() {}

func NewCasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasesContext {
	var p = new(CasesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_cases

	return p
}

func (s *CasesContext) GetParser() antlr.Parser { return s.parser }

func (s *CasesContext) COLON() antlr.TerminalNode {
	return s.GetToken(BridgeParserCOLON, 0)
}

func (s *CasesContext) CASE() antlr.TerminalNode {
	return s.GetToken(BridgeParserCASE, 0)
}

func (s *CasesContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *CasesContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(BridgeParserDEFAULT, 0)
}

func (s *CasesContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *CasesContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CasesContext) BreakStatement() IBreakStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreakStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *CasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterCases(s)
	}
}

func (s *CasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitCases(s)
	}
}

func (s *CasesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitCases(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) Cases() (localctx ICasesContext) {
	localctx = NewCasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, BridgeParserRULE_cases)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(331)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BridgeParserCASE:
		{
			p.SetState(328)
			p.Match(BridgeParserCASE)
		}
		{
			p.SetState(329)
			p.value(0)
		}

	case BridgeParserDEFAULT:
		{
			p.SetState(330)
			p.Match(BridgeParserDEFAULT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(333)
		p.Match(BridgeParserCOLON)
	}
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(334)
				p.Statement()
			}

		}
		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BridgeParserCONTINUE-32))|(1<<(BridgeParserBREAK-32))|(1<<(BridgeParserESCAPE-32)))) != 0 {
		{
			p.SetState(340)
			p.BreakStatement()
		}

	}

	return localctx
}

// ISwitchBodyContext is an interface to support dynamic dispatch.
type ISwitchBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchBodyContext differentiates from other interfaces.
	IsSwitchBodyContext()
}

type SwitchBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchBodyContext() *SwitchBodyContext {
	var p = new(SwitchBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_switchBody
	return p
}

func (*SwitchBodyContext) IsSwitchBodyContext() {}

func NewSwitchBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchBodyContext {
	var p = new(SwitchBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_switchBody

	return p
}

func (s *SwitchBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchBodyContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(BridgeParserLBRACE, 0)
}

func (s *SwitchBodyContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(BridgeParserRBRACE, 0)
}

func (s *SwitchBodyContext) AllCases() []ICasesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICasesContext)(nil)).Elem())
	var tst = make([]ICasesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICasesContext)
		}
	}

	return tst
}

func (s *SwitchBodyContext) Cases(i int) ICasesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICasesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICasesContext)
}

func (s *SwitchBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterSwitchBody(s)
	}
}

func (s *SwitchBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitSwitchBody(s)
	}
}

func (s *SwitchBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitSwitchBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) SwitchBody() (localctx ISwitchBodyContext) {
	localctx = NewSwitchBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, BridgeParserRULE_switchBody)
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
		p.SetState(343)
		p.Match(BridgeParserLBRACE)
	}
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BridgeParserCASE || _la == BridgeParserDEFAULT {
		{
			p.SetState(344)
			p.Cases()
		}

		p.SetState(349)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(350)
		p.Match(BridgeParserRBRACE)
	}

	return localctx
}

// IDeleteOpContext is an interface to support dynamic dispatch.
type IDeleteOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeleteOpContext differentiates from other interfaces.
	IsDeleteOpContext()
}

type DeleteOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteOpContext() *DeleteOpContext {
	var p = new(DeleteOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_deleteOp
	return p
}

func (*DeleteOpContext) IsDeleteOpContext() {}

func NewDeleteOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteOpContext {
	var p = new(DeleteOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_deleteOp

	return p
}

func (s *DeleteOpContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteOpContext) DELETE() antlr.TerminalNode {
	return s.GetToken(BridgeParserDELETE, 0)
}

func (s *DeleteOpContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, 0)
}

func (s *DeleteOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterDeleteOp(s)
	}
}

func (s *DeleteOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitDeleteOp(s)
	}
}

func (s *DeleteOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitDeleteOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) DeleteOp() (localctx IDeleteOpContext) {
	localctx = NewDeleteOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, BridgeParserRULE_deleteOp)

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
		p.SetState(352)
		p.Match(BridgeParserDELETE)
	}
	{
		p.SetState(353)
		p.Match(BridgeParserIdentifier)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) MemberAccess() IMemberAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberAccessContext)
}

func (s *StatementContext) FuncCall() IFuncCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *StatementContext) ExternalDecl() IExternalDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExternalDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExternalDeclContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) BreakStatement() IBreakStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreakStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *StatementContext) LoopStatement() ILoopStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopStatementContext)
}

func (s *StatementContext) SwitchStatement() ISwitchStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwitchStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) OperationWithReturn() IOperationWithReturnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationWithReturnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationWithReturnContext)
}

func (s *StatementContext) DeleteOp() IDeleteOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeleteOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeleteOpContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, BridgeParserRULE_statement)

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

	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(355)
			p.MemberAccess()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(356)
			p.FuncCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(357)
			p.ExternalDecl()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(358)
			p.Assignment()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(359)
			p.ReturnStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(360)
			p.BreakStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(361)
			p.LoopStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(362)
			p.SwitchStatement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(363)
			p.IfStatement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(364)
			p.OperationWithReturn()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(365)
			p.DeleteOp()
		}

	}

	return localctx
}

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_body
	return p
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(BridgeParserLBRACE, 0)
}

func (s *BodyContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(BridgeParserRBRACE, 0)
}

func (s *BodyContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BodyContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterBody(s)
	}
}

func (s *BodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitBody(s)
	}
}

func (s *BodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) Body() (localctx IBodyContext) {
	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, BridgeParserRULE_body)
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

	p.SetState(377)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BridgeParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(368)
			p.Match(BridgeParserLBRACE)
		}
		p.SetState(372)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BridgeParserLET)|(1<<BridgeParserSET)|(1<<BridgeParserCONST)|(1<<BridgeParserDELETE)|(1<<BridgeParserFUNC)|(1<<BridgeParserOBJ)|(1<<BridgeParserWHILE)|(1<<BridgeParserLOOP)|(1<<BridgeParserTIMES)|(1<<BridgeParserEACH)|(1<<BridgeParserIF)|(1<<BridgeParserSWITCH))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BridgeParserCONTINUE-32))|(1<<(BridgeParserRETURN-32))|(1<<(BridgeParserBREAK-32))|(1<<(BridgeParserESCAPE-32))|(1<<(BridgeParserTYPEOF-32))|(1<<(BridgeParserAT-32))|(1<<(BridgeParserNOT-32)))) != 0) || (((_la-65)&-(0x1f+1)) == 0 && ((1<<uint((_la-65)))&((1<<(BridgeParserINC-65))|(1<<(BridgeParserDEC-65))|(1<<(BridgeParserIdentifier-65)))) != 0) {
			{
				p.SetState(369)
				p.Statement()
			}

			p.SetState(374)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(375)
			p.Match(BridgeParserRBRACE)
		}

	case BridgeParserLET, BridgeParserSET, BridgeParserCONST, BridgeParserDELETE, BridgeParserFUNC, BridgeParserOBJ, BridgeParserWHILE, BridgeParserLOOP, BridgeParserTIMES, BridgeParserEACH, BridgeParserIF, BridgeParserSWITCH, BridgeParserCONTINUE, BridgeParserRETURN, BridgeParserBREAK, BridgeParserESCAPE, BridgeParserTYPEOF, BridgeParserAT, BridgeParserNOT, BridgeParserINC, BridgeParserDEC, BridgeParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(376)
			p.Statement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) FUNC() antlr.TerminalNode {
	return s.GetToken(BridgeParserFUNC, 0)
}

func (s *FunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BridgeParserLPAREN, 0)
}

func (s *FunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BridgeParserRPAREN, 0)
}

func (s *FunctionContext) AllTypeKey() []ITypeKeyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeKeyContext)(nil)).Elem())
	var tst = make([]ITypeKeyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeKeyContext)
		}
	}

	return tst
}

func (s *FunctionContext) TypeKey(i int) ITypeKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeKeyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeKeyContext)
}

func (s *FunctionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BridgeParserCOMMA)
}

func (s *FunctionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BridgeParserCOMMA, i)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, BridgeParserRULE_function)
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

	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(379)
			p.Match(BridgeParserFUNC)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(380)
			p.Match(BridgeParserFUNC)
		}
		{
			p.SetState(381)
			p.Match(BridgeParserLPAREN)
		}
		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BridgeParserFUNC)|(1<<BridgeParserOBJ)|(1<<BridgeParserINT)|(1<<BridgeParserUINT)|(1<<BridgeParserFLOAT)|(1<<BridgeParserF32)|(1<<BridgeParserF64)|(1<<BridgeParserSTRING)|(1<<BridgeParserBSTR)|(1<<BridgeParserBYTE)|(1<<BridgeParserBOOL)|(1<<BridgeParserSYMLINK)|(1<<BridgeParserANY)|(1<<BridgeParserSELF))) != 0) || _la == BridgeParserIdentifier {
			{
				p.SetState(382)
				p.TypeKey()
			}
			p.SetState(387)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == BridgeParserCOMMA {
				{
					p.SetState(383)
					p.Match(BridgeParserCOMMA)
				}
				{
					p.SetState(384)
					p.TypeKey()
				}

				p.SetState(389)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(392)
			p.Match(BridgeParserRPAREN)
		}
		p.SetState(394)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(393)
				p.TypeKey()
			}

		}

	}

	return localctx
}

// IBuiltinContext is an interface to support dynamic dispatch.
type IBuiltinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuiltinContext differentiates from other interfaces.
	IsBuiltinContext()
}

type BuiltinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltinContext() *BuiltinContext {
	var p = new(BuiltinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_builtin
	return p
}

func (*BuiltinContext) IsBuiltinContext() {}

func NewBuiltinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltinContext {
	var p = new(BuiltinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_builtin

	return p
}

func (s *BuiltinContext) GetParser() antlr.Parser { return s.parser }

func (s *BuiltinContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *BuiltinContext) SELF() antlr.TerminalNode {
	return s.GetToken(BridgeParserSELF, 0)
}

func (s *BuiltinContext) INT() antlr.TerminalNode {
	return s.GetToken(BridgeParserINT, 0)
}

func (s *BuiltinContext) UINT() antlr.TerminalNode {
	return s.GetToken(BridgeParserUINT, 0)
}

func (s *BuiltinContext) OBJ() antlr.TerminalNode {
	return s.GetToken(BridgeParserOBJ, 0)
}

func (s *BuiltinContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(BridgeParserFLOAT, 0)
}

func (s *BuiltinContext) F32() antlr.TerminalNode {
	return s.GetToken(BridgeParserF32, 0)
}

func (s *BuiltinContext) F64() antlr.TerminalNode {
	return s.GetToken(BridgeParserF64, 0)
}

func (s *BuiltinContext) STRING() antlr.TerminalNode {
	return s.GetToken(BridgeParserSTRING, 0)
}

func (s *BuiltinContext) BSTR() antlr.TerminalNode {
	return s.GetToken(BridgeParserBSTR, 0)
}

func (s *BuiltinContext) BYTE() antlr.TerminalNode {
	return s.GetToken(BridgeParserBYTE, 0)
}

func (s *BuiltinContext) BOOL() antlr.TerminalNode {
	return s.GetToken(BridgeParserBOOL, 0)
}

func (s *BuiltinContext) SYMLINK() antlr.TerminalNode {
	return s.GetToken(BridgeParserSYMLINK, 0)
}

func (s *BuiltinContext) ANY() antlr.TerminalNode {
	return s.GetToken(BridgeParserANY, 0)
}

func (s *BuiltinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuiltinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterBuiltin(s)
	}
}

func (s *BuiltinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitBuiltin(s)
	}
}

func (s *BuiltinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitBuiltin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) Builtin() (localctx IBuiltinContext) {
	localctx = NewBuiltinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, BridgeParserRULE_builtin)

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

	p.SetState(412)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BridgeParserFUNC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(398)
			p.Function()
		}

	case BridgeParserSELF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(399)
			p.Match(BridgeParserSELF)
		}

	case BridgeParserINT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(400)
			p.Match(BridgeParserINT)
		}

	case BridgeParserUINT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(401)
			p.Match(BridgeParserUINT)
		}

	case BridgeParserOBJ:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(402)
			p.Match(BridgeParserOBJ)
		}

	case BridgeParserFLOAT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(403)
			p.Match(BridgeParserFLOAT)
		}

	case BridgeParserF32:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(404)
			p.Match(BridgeParserF32)
		}

	case BridgeParserF64:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(405)
			p.Match(BridgeParserF64)
		}

	case BridgeParserSTRING:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(406)
			p.Match(BridgeParserSTRING)
		}

	case BridgeParserBSTR:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(407)
			p.Match(BridgeParserBSTR)
		}

	case BridgeParserBYTE:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(408)
			p.Match(BridgeParserBYTE)
		}

	case BridgeParserBOOL:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(409)
			p.Match(BridgeParserBOOL)
		}

	case BridgeParserSYMLINK:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(410)
			p.Match(BridgeParserSYMLINK)
		}

	case BridgeParserANY:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(411)
			p.Match(BridgeParserANY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArrayDeclContext is an interface to support dynamic dispatch.
type IArrayDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayDeclContext differentiates from other interfaces.
	IsArrayDeclContext()
}

type ArrayDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayDeclContext() *ArrayDeclContext {
	var p = new(ArrayDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_arrayDecl
	return p
}

func (*ArrayDeclContext) IsArrayDeclContext() {}

func NewArrayDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayDeclContext {
	var p = new(ArrayDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_arrayDecl

	return p
}

func (s *ArrayDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayDeclContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(BridgeParserLBRACK, 0)
}

func (s *ArrayDeclContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(BridgeParserRBRACK, 0)
}

func (s *ArrayDeclContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(BridgeParserIntegerLiteral, 0)
}

func (s *ArrayDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterArrayDecl(s)
	}
}

func (s *ArrayDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitArrayDecl(s)
	}
}

func (s *ArrayDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitArrayDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) ArrayDecl() (localctx IArrayDeclContext) {
	localctx = NewArrayDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, BridgeParserRULE_arrayDecl)
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
		p.SetState(414)
		p.Match(BridgeParserLBRACK)
	}
	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BridgeParserIntegerLiteral {
		{
			p.SetState(415)
			p.Match(BridgeParserIntegerLiteral)
		}

	}
	{
		p.SetState(418)
		p.Match(BridgeParserRBRACK)
	}

	return localctx
}

// IArrayAccessContext is an interface to support dynamic dispatch.
type IArrayAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayAccessContext differentiates from other interfaces.
	IsArrayAccessContext()
}

type ArrayAccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayAccessContext() *ArrayAccessContext {
	var p = new(ArrayAccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_arrayAccess
	return p
}

func (*ArrayAccessContext) IsArrayAccessContext() {}

func NewArrayAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayAccessContext {
	var p = new(ArrayAccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_arrayAccess

	return p
}

func (s *ArrayAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayAccessContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(BridgeParserLBRACK, 0)
}

func (s *ArrayAccessContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(BridgeParserRBRACK, 0)
}

func (s *ArrayAccessContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(BridgeParserIntegerLiteral, 0)
}

func (s *ArrayAccessContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, 0)
}

func (s *ArrayAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterArrayAccess(s)
	}
}

func (s *ArrayAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitArrayAccess(s)
	}
}

func (s *ArrayAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitArrayAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) ArrayAccess() (localctx IArrayAccessContext) {
	localctx = NewArrayAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, BridgeParserRULE_arrayAccess)
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
		p.SetState(420)
		p.Match(BridgeParserLBRACK)
	}
	{
		p.SetState(421)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BridgeParserIntegerLiteral || _la == BridgeParserIdentifier) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(422)
		p.Match(BridgeParserRBRACK)
	}

	return localctx
}

// ITypeKeyContext is an interface to support dynamic dispatch.
type ITypeKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeKeyContext differentiates from other interfaces.
	IsTypeKeyContext()
}

type TypeKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeKeyContext() *TypeKeyContext {
	var p = new(TypeKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_typeKey
	return p
}

func (*TypeKeyContext) IsTypeKeyContext() {}

func NewTypeKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeKeyContext {
	var p = new(TypeKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_typeKey

	return p
}

func (s *TypeKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeKeyContext) Builtin() IBuiltinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltinContext)
}

func (s *TypeKeyContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, 0)
}

func (s *TypeKeyContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(BridgeParserQUESTION, 0)
}

func (s *TypeKeyContext) ArrayDecl() IArrayDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayDeclContext)
}

func (s *TypeKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterTypeKey(s)
	}
}

func (s *TypeKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitTypeKey(s)
	}
}

func (s *TypeKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitTypeKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) TypeKey() (localctx ITypeKeyContext) {
	localctx = NewTypeKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, BridgeParserRULE_typeKey)

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
	p.SetState(426)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BridgeParserFUNC, BridgeParserOBJ, BridgeParserINT, BridgeParserUINT, BridgeParserFLOAT, BridgeParserF32, BridgeParserF64, BridgeParserSTRING, BridgeParserBSTR, BridgeParserBYTE, BridgeParserBOOL, BridgeParserSYMLINK, BridgeParserANY, BridgeParserSELF:
		{
			p.SetState(424)
			p.Builtin()
		}

	case BridgeParserIdentifier:
		{
			p.SetState(425)
			p.Match(BridgeParserIdentifier)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(428)
			p.Match(BridgeParserQUESTION)
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(429)
			p.ArrayDecl()
		}

	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, 0)
}

func (s *ArgumentContext) TypeKey() ITypeKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeKeyContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (s *ArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, BridgeParserRULE_argument)

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
		p.SetState(432)
		p.Match(BridgeParserIdentifier)
	}
	{
		p.SetState(433)
		p.TypeKey()
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, 0)
}

func (s *ArgumentsContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(BridgeParserELLIPSIS, 0)
}

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BridgeParserCOMMA)
}

func (s *ArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BridgeParserCOMMA, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, BridgeParserRULE_arguments)
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

	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(435)
			p.Match(BridgeParserIdentifier)
		}
		{
			p.SetState(436)
			p.Match(BridgeParserELLIPSIS)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(437)
			p.Argument()
		}
		p.SetState(442)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == BridgeParserCOMMA {
			{
				p.SetState(438)
				p.Match(BridgeParserCOMMA)
			}
			{
				p.SetState(439)
				p.Argument()
			}

			p.SetState(444)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IFuncDeclContext is an interface to support dynamic dispatch.
type IFuncDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncDeclContext differentiates from other interfaces.
	IsFuncDeclContext()
}

type FuncDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclContext() *FuncDeclContext {
	var p = new(FuncDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_funcDecl
	return p
}

func (*FuncDeclContext) IsFuncDeclContext() {}

func NewFuncDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclContext {
	var p = new(FuncDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_funcDecl

	return p
}

func (s *FuncDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclContext) FUNC() antlr.TerminalNode {
	return s.GetToken(BridgeParserFUNC, 0)
}

func (s *FuncDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BridgeParserLPAREN, 0)
}

func (s *FuncDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BridgeParserRPAREN, 0)
}

func (s *FuncDeclContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *FuncDeclContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, 0)
}

func (s *FuncDeclContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FuncDeclContext) TypeKey() ITypeKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeKeyContext)
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterFuncDecl(s)
	}
}

func (s *FuncDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitFuncDecl(s)
	}
}

func (s *FuncDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitFuncDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) FuncDecl() (localctx IFuncDeclContext) {
	localctx = NewFuncDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, BridgeParserRULE_funcDecl)
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
		p.SetState(447)
		p.Match(BridgeParserFUNC)
	}
	p.SetState(449)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BridgeParserIdentifier {
		{
			p.SetState(448)
			p.Match(BridgeParserIdentifier)
		}

	}
	{
		p.SetState(451)
		p.Match(BridgeParserLPAREN)
	}
	p.SetState(453)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BridgeParserIdentifier {
		{
			p.SetState(452)
			p.Arguments()
		}

	}
	{
		p.SetState(455)
		p.Match(BridgeParserRPAREN)
	}
	p.SetState(457)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(456)
			p.TypeKey()
		}

	}
	{
		p.SetState(459)
		p.Body()
	}

	return localctx
}

// IDeclaratorContext is an interface to support dynamic dispatch.
type IDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclaratorContext differentiates from other interfaces.
	IsDeclaratorContext()
}

type DeclaratorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaratorContext() *DeclaratorContext {
	var p = new(DeclaratorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_declarator
	return p
}

func (*DeclaratorContext) IsDeclaratorContext() {}

func NewDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaratorContext {
	var p = new(DeclaratorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_declarator

	return p
}

func (s *DeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaratorContext) CONST() antlr.TerminalNode {
	return s.GetToken(BridgeParserCONST, 0)
}

func (s *DeclaratorContext) LET() antlr.TerminalNode {
	return s.GetToken(BridgeParserLET, 0)
}

func (s *DeclaratorContext) SET() antlr.TerminalNode {
	return s.GetToken(BridgeParserSET, 0)
}

func (s *DeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterDeclarator(s)
	}
}

func (s *DeclaratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitDeclarator(s)
	}
}

func (s *DeclaratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitDeclarator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) Declarator() (localctx IDeclaratorContext) {
	localctx = NewDeclaratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, BridgeParserRULE_declarator)
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
		p.SetState(461)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BridgeParserLET)|(1<<BridgeParserSET)|(1<<BridgeParserCONST))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_decl
	return p
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) Declarator() IDeclaratorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaratorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaratorContext)
}

func (s *DeclContext) Identifiers() IIdentifiersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifiersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifiersContext)
}

func (s *DeclContext) TypeKey() ITypeKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeKeyContext)
}

func (s *DeclContext) AssignmentOp() IAssignmentOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentOpContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitDecl(s)
	}
}

func (s *DeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) Decl() (localctx IDeclContext) {
	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, BridgeParserRULE_decl)

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
		p.SetState(463)
		p.Declarator()
	}
	{
		p.SetState(464)
		p.Identifiers()
	}
	p.SetState(466)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(465)
			p.TypeKey()
		}

	}
	p.SetState(469)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(468)
			p.AssignmentOp()
		}

	}

	return localctx
}

// IClassBodyContext is an interface to support dynamic dispatch.
type IClassBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassBodyContext differentiates from other interfaces.
	IsClassBodyContext()
}

type ClassBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassBodyContext() *ClassBodyContext {
	var p = new(ClassBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_classBody
	return p
}

func (*ClassBodyContext) IsClassBodyContext() {}

func NewClassBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassBodyContext {
	var p = new(ClassBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_classBody

	return p
}

func (s *ClassBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassBodyContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(BridgeParserLBRACE, 0)
}

func (s *ClassBodyContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(BridgeParserRBRACE, 0)
}

func (s *ClassBodyContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *ClassBodyContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ClassBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterClassBody(s)
	}
}

func (s *ClassBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitClassBody(s)
	}
}

func (s *ClassBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitClassBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) ClassBody() (localctx IClassBodyContext) {
	localctx = NewClassBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, BridgeParserRULE_classBody)
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
		p.SetState(471)
		p.Match(BridgeParserLBRACE)
	}
	p.SetState(475)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BridgeParserIdentifier {
		{
			p.SetState(472)
			p.Argument()
		}

		p.SetState(477)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(478)
		p.Match(BridgeParserRBRACE)
	}

	return localctx
}

// IClassDeclContext is an interface to support dynamic dispatch.
type IClassDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassDeclContext differentiates from other interfaces.
	IsClassDeclContext()
}

type ClassDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDeclContext() *ClassDeclContext {
	var p = new(ClassDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BridgeParserRULE_classDecl
	return p
}

func (*ClassDeclContext) IsClassDeclContext() {}

func NewClassDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDeclContext {
	var p = new(ClassDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BridgeParserRULE_classDecl

	return p
}

func (s *ClassDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDeclContext) OBJ() antlr.TerminalNode {
	return s.GetToken(BridgeParserOBJ, 0)
}

func (s *ClassDeclContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BridgeParserIdentifier, 0)
}

func (s *ClassDeclContext) ClassBody() IClassBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassBodyContext)
}

func (s *ClassDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.EnterClassDecl(s)
	}
}

func (s *ClassDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BridgeListener); ok {
		listenerT.ExitClassDecl(s)
	}
}

func (s *ClassDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BridgeVisitor:
		return t.VisitClassDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BridgeParser) ClassDecl() (localctx IClassDeclContext) {
	localctx = NewClassDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, BridgeParserRULE_classDecl)

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
		p.SetState(480)
		p.Match(BridgeParserOBJ)
	}
	{
		p.SetState(481)
		p.Match(BridgeParserIdentifier)
	}
	{
		p.SetState(482)
		p.ClassBody()
	}

	return localctx
}

func (p *BridgeParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 8:
		var t *ValueContext = nil
		if localctx != nil {
			t = localctx.(*ValueContext)
		}
		return p.Value_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *BridgeParser) Value_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 14)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
