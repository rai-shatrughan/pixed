import React from 'react';
import { createRoot } from "react-dom/client";
import About from './about/About';

const container = document.getElementById("root");
const root = createRoot(container);
root.render(<About />);