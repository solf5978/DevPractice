for (int i = 0; i < n; i++)
{
    A[i] = i;
}
for (int i = 0; i < m; i++)
{
    scanf("%s%d%d", type, &X, &Y);
    p = find(X);
    q = find(Y);
    if (type[0] == 'A')
    {
        if (q > p)
        {
            shift_circular_left(p, q - 1);
        }
        else
            shift_circular_right(q, p);
    }
    else
    {
        if (q > p)
        {
            shift_circular_left(p, q);
        }
        else
            shift_circular_right(q + 1, p);        
    }
}

