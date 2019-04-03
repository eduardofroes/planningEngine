#include<stdlib.h>
#include<stdio.h>
#include<strings.h>


int size = 0;

typedef struct rule{
    char name[50];
    char value[50];
} Rule;

Rule* newRule(char *name, char *value){
    Rule *rule = (Rule *) malloc(sizeof(Rule));

    if(rule != NULL){
        *rule->name = *name;
        *rule->value = *value;
    } else {
        printf("Error in memory allocation.");
    }   
    return(rule);
}

Rule* appendRule(Rule* rules, Rule* rule) {
    Rule *temp_rules = (Rule *) realloc(rules, (size+1)*sizeof(Rule));        
     
    if(temp_rules != NULL){
        *(temp_rules+size) = *rule;
        size++;
    } else {
        printf("Error in memory allocation.");
    }
    
    return temp_rules;
}