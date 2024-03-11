export namespace todo {
	
	export class List {
	    id: number;
	    title: string;
	
	    static createFrom(source: any = {}) {
	        return new List(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	    }
	}
	export class Task {
	    id: number;
	    list_id: number;
	    title: string;
	    is_done: boolean;
	    starred: boolean;
	    description: string;
	    deadline: string;
	    reward: number;
	    elapsed: number;
	    created_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Task(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.list_id = source["list_id"];
	        this.title = source["title"];
	        this.is_done = source["is_done"];
	        this.starred = source["starred"];
	        this.description = source["description"];
	        this.deadline = source["deadline"];
	        this.reward = source["reward"];
	        this.elapsed = source["elapsed"];
	        this.created_at = source["created_at"];
	    }
	}

}

