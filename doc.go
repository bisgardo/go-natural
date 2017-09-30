/*
	Package strcmp is a package containing a comparator function
	and a companion sort.Interface implementation
	for sorting strings according to "natural order".

	The precise definition of natural comparison is the following:

	Strings are equal if and only if they consist of the same bytes in the same order. If they are different, the order is determined as follows:

	Write the compared strings left and right as the longest common prefix followed by (non-empty) suffixes:

	left  = <prefix><suffix1>
	right = <prefix><suffix2>

	Extract any prefix number from the suffixes into a separate group, such that none of the suffixes start with a number character:

	left  = <prefix><number1><suffix1>
	right = <prefix><number2><suffix2>
	If both number groups are non-empty, the string with the greater number is "larger". If the numbers are equal but have a different number of leading zeros, the number with the most leading zeros is "larger".

	If either of the number groups are empty, the order is simply defined according to the usual definition (i.e. byte comparison).
*/
package strcmp
