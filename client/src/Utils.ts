export interface ParamDef {
    name: string;
    type?: string;
}
export interface PropDef {
    name: string;
    type?: string;
    required: boolean;
    defaultValue?: string;
    description?: string;
    returnType?: string;
    params?: Array<ParamDef>;
}

export interface ComponentDef {
    name: string;
    sourcePath: string;
    componentType: string;
    description: string;
    props?: Array<PropDef>
}

export function componentSorter(a: ComponentDef, b: ComponentDef) {
    return a.name.localeCompare(b.name);
}
