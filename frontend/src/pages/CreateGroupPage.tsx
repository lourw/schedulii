import type { ReactElement } from "react";
import React from "react";
import GroupForm from "../components/GroupForm";
import Navbar from "../components/Navbar";

const CreateGroupPage = (): ReactElement => (
    <div>
        <Navbar />
        <GroupForm />
    </div>
);

export default CreateGroupPage;
