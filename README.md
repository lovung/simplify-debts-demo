# Problems

Read my blog post [here](https://blog.vulong.dev/Simplify-Debts-algorithm-a528f72ff6f44a16a62b78e36d1a6268)

*What does `Splitwise`’s ‘simplify debts’ feature do?*

> Simplify debts (a.k.a. “debt simplification”) is a feature of `Splitwise` that restructures debt within groups of people. 
It does not change the total amount that anyone owes, but it makes it easier to pay people back by minimizing the total number of payments.
> 

> For example: say, Anna, Bob, and Charlie share an apartment. Anna owes Bob $20, and Bob owes Charlie $20. Rather than making two separate payments, `Splitwise` would tell Anna to pay $20 to Charlie directly, thereby minimizing the total number of payments being made. This ensures that people are paid back more quickly and efficiently.’
> 

Consider a group of seven people namely Alice, Bob, Charlie, David, Ema, Fred and Gabe. They went out for a tour together and at the end of the tour realized that they have the following debts;

```bash
Gabe owes $30 to Bob.
Gabe owes $10 to David.
Fred owes $10 to Bob.
Fred owes $30 to Charlie.
Fred owes $10 to David.
Fred owes $10 to Ema.
Bob owes $40 to Charlie.
Charlie owes $20 to David.
David owes $50 to Ema.

```

![**Figure 1. Representing Debts in the form of a Directed Graph**](https://user-images.githubusercontent.com/17726269/185643245-804bbbd6-d898-409a-a2c2-b7c732c576db.png)

**Figure 1. Representing Debts in the form of a Directed Graph**

# There is an article on Medium

[https://medium.com/@mithunmk93/algorithm-behind-splitwises-debt-simplification-feature-8ac485e97688](https://medium.com/@mithunmk93/algorithm-behind-splitwises-debt-simplification-feature-8ac485e97688)

## First step

```bash
**Net Change in Cash = (Sum of Cash Inflow - Sum of Cash Outflow)**
```

It can be noted that the total money owed by *Givers* is always equal to the total money to be received by *Receivers*

![**Figure 2. Clustering Givers and Receivers into two different groups**](https://user-images.githubusercontent.com/17726269/185643443-1877610c-1c69-4f18-8265-88cb155cec9a.png)

**Figure 2. Clustering Givers and Receivers into two different groups**

## But the result in this article is not optimized

![**Figure 3. Simplified debts graph returned by the algorithm**](https://user-images.githubusercontent.com/17726269/185643622-e8a2c891-1673-4d1d-a892-ebfb1f7987e8.png)

**Figure 3. Simplified debts graph returned by the algorithm**


![**Figure 4. The optimized result should be** ](https://user-images.githubusercontent.com/17726269/185643711-fd95d62d-3b4f-44d9-8904-dfeff3dfa3ae.png)

**Figure 4. The optimized result should be**

# Deep diving

## V vertices, max V **- 1 edge in simplified prove**

### **Adding and Subtracting Vectors**

![image](https://user-images.githubusercontent.com/17726269/185644152-bf5781c0-8a06-4950-8c78-698e1cac8212.png)

**Parallelogram Method**

![image](https://user-images.githubusercontent.com/17726269/185644232-5daade75-c8b0-4243-b918-b531e0922f1a.png)
![image](https://user-images.githubusercontent.com/17726269/185644261-b70b4ac1-7977-4fd4-9127-7d1cd0f2a837.png)

**Triangle Method**

![image](https://user-images.githubusercontent.com/17726269/185644326-48b03de4-df10-4efe-982c-c10d430d8030.png)
![image](https://user-images.githubusercontent.com/17726269/185644371-8c9d9694-a9b7-4f4b-aaac-ff93226e13c0.png)

### Apply for this problem

For 3 vertices, we have maximum 6 unidirectional edges: A → B, B → A, A → C, C → A, B → C, C → B with positive value.

$\overrightarrow{AB}$

We can simplify A → B, B → A to $\overrightarrow{AB}$; B → C, C → B to $\overrightarrow{BC}$; and A → C, C → A to $\overrightarrow{CA}$

![image](https://user-images.githubusercontent.com/17726269/185644457-ac0305e6-ebbb-4ae9-ac22-44de700e10e6.png)

So, we can convert this situation by splitting $\overrightarrow{CA}$ to $\overrightarrow{AB}$ and  $\overrightarrow{BC}$

![image](https://user-images.githubusercontent.com/17726269/185644517-b20d8cce-94ee-43a9-a4ab-4d1dff17d06d.png)

When we have the new vertex comes, we can also simplify the vector with current nodes into 1 vector:

![image](https://user-images.githubusercontent.com/17726269/185644562-55341f2e-4c61-4f8d-991a-943db3b91427.png)

Continuously, for V vertices, we can simplify to V - 1 edges.

## The minimum edge in one group

For a big `zero-sum` group with `V` vertices, if we can split it into `k` smaller `zero-sum` groups.

As we’ve already proven above, with the big group, the number of internal optimized edges is `N - 1`.

| Number of big group | Number of internal optimized edges |
| --- | --- |
| 1 | N - 1 |
| 2 | (N1 - 1) + (N2 - 1) = (N1 + N2) - 2 = N - 2 |
| 3 | N - 3 |
| ... | ... |
| k | N - k |

But how?

→ This is Subset Sum Problem ([https://en.wikipedia.org/wiki/Subset_sum_problem](https://en.wikipedia.org/wiki/Subset_sum_problem)) and it is [NP-hard](https://en.wikipedia.org/wiki/NP-hardness):

> if there are `n` integers in the `nums` list, there exist `2^n — 1`
subsets that need to be checked (excluding the empty set).
> 

We can use the Dynamic Programming (DP) to solve it.

But 

> **DP has shortcomings, brute-force can be better 💪**
Let’s says `nums = [-1000000, 1000000]` and `target = 999999`. Using the DP method, we would have 2 rows and `999999 + 1000000 + 1 = 2000000` columns. That’s a lot of memory usage for an obviously unsolvable problem! If there are few numbers in `nums` but the range of values is large, you’re better off brute-force checking all possible subsets.
> 

## But is the optimal result is always good for your product? Why the result above is not optimal?

Let reach back to me with a reference to a *[Splitwise page](https://blog.splitwise.com/2012/09/14/debts-made-simple/)* that mentioned about 3 rules the feature obeys, which are as listed below:

1. *Everyone owes the same net amount at the end,*
2. ***No one owes a person that they didn’t owe before, and***
3. *No one owes more money in total than they did before the simplification.*

So the problem boils down to ****varying the amount being transferred on existing transactions without introducing newer ones**.** This algorithmically translates to the following,

> ***Given a Directed Graph representing Debts (as shown in Figure 1), change (if needed) the weights on the existing edges without introducing newer ones***.
> 

They want to keep the old `debts` aka `edges` to apply the [Maximum-Flow Algorithm](https://en.wikipedia.org/wiki/Maximum_flow_problem).

For example, in the `mithunmk93`'s article, he use [Dinic's_algorithm](https://en.wikipedia.org/wiki/Dinic%27s_algorithm).

But the second rule seems the problem here.

![image](https://user-images.githubusercontent.com/17726269/185644739-19fbdd13-8b8d-43ae-b5bf-6a2ceac3080d.png)

A flow network, with source *s* and sink *t*. The numbers next to the edge are the capacities.

# My solution

[https://github.com/lovung/simplify-debts-demo](https://github.com/lovung/simplify-debts-demo)

## My rules

1. *Everyone owes the same net amount at the end.*
2. *Good perform for < 20 members (vertices).*
3. *Only Givers need to make the transfer, and the all destinations should be a Receiver (should not transfer to other Giver).*
4. (Optional) *Minimum the numbers of edges as much as possible (no need to optimal).*
5. (Optional) *Minimize the maximum edge of 1 Giver-vertex as much as possible. Receivers don’t care about the incoming edges because they don’t make the transfer.*

## Steps

1. (Optional) We need to find out as many as possible `zero-sum` pairs by using hash map. Create edges inside the pair.
    - Time complexity `O(V)`
        
        Time complexity if want to find out `3-vertices` subset: $O({V^2 \above{1pt} 4})$ because ${V^2 \above{1pt} 4} ≥  {(a+b)^2 \above{1pt} 4} ≥ ab$.
        
2. Remove all vertices which already grouped in Step 1.
3. Sort the positive numbers (Receivers) and negative numbers (Givers) by the absolute value (ASC).
4. (Optional) For each Receiver, find the smallest-least-recently-transfer Giver who still owes > Receiver’s amount.
    - Why is it optional?
        
        Earlier, I want to support the 5th rule by using this step. 
        
        But seem it’s not working in test case 4.x and conflict to 4th rule.
        
5. Map the remain Receivers and Givers until their balances is zero.

![image](https://user-images.githubusercontent.com/17726269/185644791-d05a7d62-3c41-4a6a-a64d-0a6a9ba22482.png)

Without smallest-least-recently-transfer

![image](https://user-images.githubusercontent.com/17726269/185644833-00a6f5a8-dbe2-4f81-9ac2-8d4588dd1f8f.png)

With smallest-least-recently-transfer

## Results

Receivers: 1, 2, 3, 5, 6, 7, 10, 11

Givers: -1, -4, -6, -7, -7, -9, -10

| Test case 1 | Step 1: true | Step 1: false |
| --- | --- | --- |
| Step 4: true | (case 1.0) 11 edges | (case 1.2) 14 edges |
| Step 4: false | (case 1.1) 10 edges | (case 1.3) 12 edges |

## Benchmarks

Receivers: 1, 2, 3, 5, 6, 7, 10, 11

Givers: -1, -4, -6, -7, -7, -9, -10

```bash
BenchmarkTwoSidesGraph_1_0-12    	  171507	      7202 ns/op	    6579 B/op	     116 allocs/op
BenchmarkTwoSidesGraph_1_1-12    	  176533	      7150 ns/op	    6547 B/op	     115 allocs/op
BenchmarkTwoSidesGraph_1_2-12    	  193450	      5467 ns/op	    5768 B/op	      89 allocs/op
BenchmarkTwoSidesGraph_1_3-12    	  241360	      5639 ns/op	    5704 B/op	      87 allocs/op
```

# References

1. Subset sum problem:
    1. [https://towardsdatascience.com/how-to-find-all-solutions-to-the-subset-sum-problem-597f77677e45](https://towardsdatascience.com/how-to-find-all-solutions-to-the-subset-sum-problem-597f77677e45)
    2. [https://en.wikipedia.org/wiki/Subset_sum_problem](https://en.wikipedia.org/wiki/Subset_sum_problem)
2. Knapsack Problem:
    1. [https://medium.com/@fabianterh/how-to-solve-the-knapsack-problem-with-dynamic-programming-eb88c706d3cf](https://medium.com/@fabianterh/how-to-solve-the-knapsack-problem-with-dynamic-programming-eb88c706d3cf)
    2. [https://medium.com/@fabianterh/optimizing-the-knapsack-problem-dynamic-programming-solution-for-space-complexity-c6bcdff3870b](https://medium.com/@fabianterh/optimizing-the-knapsack-problem-dynamic-programming-solution-for-space-complexity-c6bcdff3870b)
