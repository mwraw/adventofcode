#include <stdio.h>
#include <stdlib.h>

typedef struct node {
	int marble;
	struct node *previous;
	struct node *next;
} node_t;

node_t *add_to_list(long, node_t *);
node_t *move_in_list(node_t *, int, int);
node_t *delete_from_list(node_t *);
long winning_score(long *, int);

int main(int argc, char *argv[]) {
	node_t *current = malloc(sizeof(node_t));
	
	if(argc != 3) {
		printf("Wrong number of arguments!\n");
		exit(0);
	}
	
	int num_players = strtod(argv[1], 0);
	long highest_marble = strtol(argv[2], 0, 10);
	
	long *scores = calloc(num_players, sizeof(long));
	
    current->marble = 0;
    current->next = current;
    current->previous = current;
	
	for(long marble = 1; marble<=highest_marble; marble++) {
		if(marble % 23 == 0) {
			current = move_in_list(current, 7, -1);
			scores[(marble-1)%num_players] += marble + current->marble;
			current = delete_from_list(current);
		} else {
			current = move_in_list(current, 1, 1);
			current = add_to_list(marble, current);
		}
	}
	
	printf("%ld\n", winning_score(scores, num_players));
	
	return 0;
}

long winning_score(long *scores, int num_players) {
	long winning_score = 0;
	
	for(int player = 0; player<num_players; player++)
		if(scores[player]>winning_score)
			winning_score = scores[player];
	
	return winning_score;
}

node_t *add_to_list(long marble, node_t *after) {
	node_t *new_marble = malloc(sizeof(node_t));

	new_marble->marble = marble;
    new_marble->previous = after;
    new_marble->next = after->next;

    after->next->previous = new_marble;
    after->next = new_marble;

    return new_marble;
}

node_t *move_in_list(node_t *current, int steps, int direction) {
    for(; steps > 0; steps--)
		current = direction == 1 ? current->next : current->previous;
		
	return current;
}

node_t *delete_from_list(node_t *current_marble) {
	current_marble->previous->next = current_marble->next;
    current_marble->next->previous = current_marble->previous;
	
    return current_marble->next;
}