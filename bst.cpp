#include <bits/stdc++.h>
using namespace std;

/* A binary tree node has data,
pointer to left child and a
pointer to right child */
struct Node
{
    int data;
    Node *left, *right;
};

/* Helper function that allocates a
new node */
Node *insert(Node *root, int data)
{
    if (root == NULL)
    {
        Node *temp = new Node;
        temp->data = data;
        temp->left = temp->right = NULL;
        return temp;
    }
    else if (data <= root->data)
        root->left = insert(root->left, data);
    else
        root->right = insert(root->right, data);
    return root;
}

Node *insertInOrder(int arr[])
{
    Node *root = NULL;
    for (int i = 0; i < sizeof(arr) / sizeof(arr[0]); i++)
    {
        root = insert(root, arr[i]);
    }
    return root;
}

// function display the tree in inorder
void inOrder(Node *root)
{
    if (root == NULL)
        return;
    inOrder(root->left);
    cout << root->data << " ";
    inOrder(root->right);
}

int main()
{
    int arr[] = {10, 6, 13, 4, 8, 12, 15};
    // for loop to input array elements
    for (int i = 0; i < sizeof(arr) / sizeof(arr[0]); i++)
    {
        cout << arr[i] << " ";
    }
    int n = sizeof(arr) / sizeof(arr[0]);
    Node *root = insertInOrder(arr);
    inOrder(root);
}