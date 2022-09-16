import { createSlice } from '@reduxjs/toolkit';

const initialState = {
    value: 0,
}

const counterSlice = createSlice({
    name: 'counter',
    initialState,
    reducers: {
        // imer lib makes this immutability possible
        // increment
        increment : (state) => {
            state.value++;
        },
        // decrement
        decrement : (state) => {
            state.value--;
        }
    }
})

export const { increment, decrement } = counterSlice.actions;

export default counterSlice.reducer;