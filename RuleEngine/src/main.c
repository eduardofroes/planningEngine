#include <stdio.h>
#include <stdlib.h>
#include "Rule.h"
#include <string.h>

int main() {
	char name[50], value[50];
	Rule *rules = NULL;

	for(int i = 0; i < 3; i++)
	{
		printf("\nName:");
		scanf("%s", name);

		printf("\nValue:");
		scanf("%s", value);

		Rule *rule = newRule(name, value);
		rules = appendRule(rules, rule);

		printf("\nINDEX: %i -> Rule->Name:%s\tRule->Value:%s\n", i, rules->name, rules->value);
	}

	Rule *temp = rules;

	for(int i = 0; i < 3; i++)
	{
		printf("\nINDEX: %i -> Rule->Name:%s\tRule->Value:%s\n", i, temp->name, temp->value);
		temp++;
	}
	

	printf("\nCount:%f", (float)(sizeof(rules)/sizeof(Rule)));

	return(0);
}