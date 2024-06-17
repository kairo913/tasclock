export interface List {
    Id: number;
    Title: string;
    Description: string;
    CreatedAt: string;
    UpdatedAt: string;
}

export interface Task {
    Id: number;
    ListId: number;
    Title: string;
    Description: string;
    Completed: boolean;
    CreatedAt: string;
    UpdatedAt: string;
}