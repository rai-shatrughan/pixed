import torch

x = torch.arange(12, dtype=torch.float32)
y = x.reshape(3, 4)
t0 = torch.zeros((2, 3, 4))


# print(x, "\n")
# print("num of el - ", x.numel())
# print("size - ", x.size())
# print("reshape - ", y)
print("t0 - ", t0)
