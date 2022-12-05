import torch

x = torch.arange(12, dtype=torch.float32)
X = x.reshape(3, 4)
t0 = torch.zeros((2, 3, 4))
t1 = torch.ones((3, 2, 4))
tr = torch.randn((2, 3, 4))
ta = torch.tensor([[2, 1, 4, 3], [1, 2, 3, 4], [4, 3, 2, 1]])
X[:2, :] = 12

# print(x, "\n")
# print("num of el - ", x.numel())
# print("size - ", x.size())
print("reshape - ", X)
# print("t0 - ", t0)
# print("t1 - ", t1)
# print("tr - ", tr)
# print("ta - ", ta)
# print("X[-1] - ", X[-1])
# print("X[1:3] - ", X[1:3])
