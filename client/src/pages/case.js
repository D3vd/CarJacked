import React from "react";
import { Router } from "@reach/router";

import PublicCase from "../components/PublicCase";

const Case = () => (
  <Router>
    <PublicCase path="/case/:id" />
  </Router>
);

export default Case;
