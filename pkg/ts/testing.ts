

import { GQLClient, GQLResponse, buildQuery } from "./package/common"


export interface Group {
    name:string;
    attributes:string[];
    
}

export interface GroupSelection {
    [index: string]: any;
    name?:boolean;
    attributes?:boolean |  {limit:string;};
    
}

export interface User {
    isActive:boolean;
    firstName:string;
    groups:number[];
    mainGroup:Group;
    
}

export interface UserSelection {
    [index: string]: any;
    isActive?:boolean;
    firstName?:boolean;
    groups?:boolean;
    mainGroup?:{name?:boolean;attributes?:boolean;};
    
}


export async function queryUsers({sort, filter, selections}: {sort:string, filter?:string, selections: UserSelection}): Promise<GQLResponse<User>> {
    let query = buildQuery("users", selections);
    const client = new GQLClient();
    const response: GQLResponse<User> = await client.post<User>("", {sort:sort, filter:filter});
    return response;
}


// early testing...
const users = await queryUsers({ sort: "", selections: { isActive: true, mainGroup: { name: true } }})