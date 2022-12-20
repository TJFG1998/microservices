#include <stdio.h>
#include <string.h>
#include <stdlib.h>


struct UserData { 
	int nr_participant;
	char *name;
	int cellphone_number;
	int age;
    struct UserData *next;
    struct UserData *prev;
} UserData;


void printList(struct UserData* users_list)
{
    struct UserData *temp = users_list;
    while (temp != NULL) {
        printf(" %d ", temp->nr_participant);
        temp = temp->next;
    }
}

FILE* fetchFile(char *file_path)
{
    FILE *arq = fopen(file_path, "rt");

	if(arq == NULL)
	{
		printf("FILE PROVIDED DOES NOT EXIST");
	}

    return arq;
}

void insertUserDataOnList(char *file_line, struct UserData **users_list)
{
    char *splited_string = strtok(file_line, ";");
    char *data[4];

    int i = 0;

    struct UserData *new_node = NULL;
    new_node = (struct UserData*)malloc(sizeof(struct UserData));
    new_node->next = NULL;
    new_node->prev = NULL;

    while(splited_string != NULL)
    {
        data[i++] = splited_string;
        splited_string = strtok(NULL, ";");
    }

    new_node->nr_participant = atoi(data[0]);
    new_node->name = data[1];
    new_node->cellphone_number = atoi(data[2]);
    new_node->age = atoi(data[3]);

    if(*users_list == NULL)
    {
        *users_list = new_node;
    }
    else
    {
        struct UserData *last_node = *users_list;
        
        while(last_node->next != NULL)
        {
            last_node = last_node->next;
        }
        last_node->next = new_node;
        new_node->prev = last_node;
    }
}

void readUsersData(FILE *file, struct UserData* users_list)
{
	char *line;
	char *result;
    while(!feof(file))
	{
		result = fgets(line, 510, file);
		if (result)
        {
            printf("1\n");
            insertUserDataOnList(line, &users_list);
        }
	}
}

int main(int argc, char *argv[])
{
    FILE *arq;
    struct UserData *users_list = NULL;

//	if(argc != 2)
//	{
//		printf("NO FILE PATH PROVIDED\n");
//		return 1;
//	}

	arq = fetchFile("/home/shangming/workspace/c_go_back/users_1.txt");

    readUsersData(arq, users_list);

    printList(users_list);

	fclose(arq);
	return 0;
}
