package main

var testsbnpshareinput = [3]string{
	"3\n10 20 30\n",
	"4\n10 20 30 40\n",
	"4\n13 15 18 20\n"}

var testsbnpbndinput = [3]string{
	"6\n1 5\n 5 10\n10 10\n20 20\n50 5\n100 10\n",
	"3\n1 20\n72 50\n100 30\n",
	"3\n2.56 20\n3.72 20\n4.55 26\n"}

var testsbnp = [3]BankNoteProblem{
	{
		[]int{
			10,
			20,
			30},
		[]*BankNoteDeck{
			{1, 5},
			{5, 10},
			{10, 10},
			{20, 20},
			{50, 5},
			{100, 10}}},
	{
		[]int{
			10,
			20,
			30,
			40},
		[]*BankNoteDeck{
			{1, 20},
			{72, 50},
			{100, 30}}},
	{
		[]int{
			13,
			15,
			18,
			20},
		[]*BankNoteDeck{
			{2.56, 20},
			{3.72, 20},
			{4.55, 26}}}}

var testsbns = [3]BankNoteSolution{

	{[]*RobberAccount{
		{
			[]*BankNoteDeck{
				{1, 1},
				{5, 0},
				{10, 5},
				{20, 0},
				{50, 3},
				{100, 1}}},
		{
			[]*BankNoteDeck{
				{1, 2},
				{5, 8},
				{10, 4},
				{20, 1},
				{50, 0},
				{100, 5}}},
		{
			[]*BankNoteDeck{
				{1, 2},
				{5, 2},
				{10, 1},
				{20, 19},
				{50, 2},
				{100, 4}}}}},
	{[]*RobberAccount{
		{
			[]*BankNoteDeck{
				{1, 2},
				{72, 5},
				{100, 3}}},
		{
			[]*BankNoteDeck{
				{1, 4},
				{72, 10},
				{100, 6}}},
		{
			[]*BankNoteDeck{
				{1, 6},
				{72, 15},
				{100, 9}}},
		{
			[]*BankNoteDeck{
				{1, 8},
				{72, 20},
				{100, 12}}}}},
	{[]*RobberAccount{
		{
			[]*BankNoteDeck{
				{2.56, 1},
				{3.72, 11},
				{4.55, 1}}},
		{
			[]*BankNoteDeck{
				{2.56, 6},
				{3.72, 1},
				{4.55, 8}}},
		{
			[]*BankNoteDeck{
				{2.56, 6},
				{3.72, 4},
				{4.55, 8}}},
		{
			[]*BankNoteDeck{
				{2.56, 7},
				{3.72, 4},
				{4.55, 9}}}}}}
