import pandas as pd
import numpy as np
import matplotlib.pyplot as plt

print("pd version", pd.__version__, "\n")
print("np version", np.__version__, "\n")

series = pd.Series([20, 21, 31, 42, 47, 54, 94])

print("series - ", series, "\n")

df = pd.DataFrame(
    [[9, 6, 8],
    [12, 40, 60],
    [53, 86, 43]],
    index=[1, 2, 3],
    columns=['a', 'b', 'c']
)

print(df,"\n")

x = np.linspace(0, 2 * np.pi, 200)
y = np.sin(x)

fig, ax = plt.subplots()
ax.plot(x, y)
plt.show()