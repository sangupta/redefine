export interface NoProps {

}

export interface ParamDef {
    name: string;
    type?: string;
}

export interface PropDef {
    name: string;
    type?: string;
    required: boolean;
    enumOf?: Array<ParamDef>
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
    docs: string;
}

export function componentSorter(a: ComponentDef, b: ComponentDef) {
    if (!a || !b) {
        console.warn('sorting received undefined');
        return 0;
    }

    return a.name.localeCompare(b.name);
}

export function propsSorter(a: PropDef, b: PropDef) {
    if (!a || !b) {
        console.warn('sorting received undefined');
        return 0;
    }

    return a.name.localeCompare(b.name);
}