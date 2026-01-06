export namespace terminal {
	
	export class CommandEntry {
	    id: string;
	    command: string;
	    output: string;
	    exitCode: number;
	    startTime: string;
	    endTime: string;
	    duration: number;
	
	    static createFrom(source: any = {}) {
	        return new CommandEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.command = source["command"];
	        this.output = source["output"];
	        this.exitCode = source["exitCode"];
	        this.startTime = source["startTime"];
	        this.endTime = source["endTime"];
	        this.duration = source["duration"];
	    }
	}

}

export namespace report {
	
	export class Finding {
	    id: string;
	    title: string;
	    severity: string;
	    description: string;
	    evidence: string;
	    command: string;
	    timestamp: string;
	    remediation: string;
	
	    static createFrom(source: any = {}) {
	        return new Finding(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.severity = source["severity"];
	        this.description = source["description"];
	        this.evidence = source["evidence"];
	        this.command = source["command"];
	        this.timestamp = source["timestamp"];
	        this.remediation = source["remediation"];
	    }
	}
	export class Report {
	    id: string;
	    title: string;
	    startTime: string;
	    endTime: string;
	    findings: Finding[];
	    commandTimeline: terminal.CommandEntry[];
	    summary: string;
	    methodology: string;
	    recommendations: string[];
	
	    static createFrom(source: any = {}) {
	        return new Report(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.startTime = source["startTime"];
	        this.endTime = source["endTime"];
	        this.findings = this.convertValues(source["findings"], Finding);
	        this.commandTimeline = this.convertValues(source["commandTimeline"], terminal.CommandEntry);
	        this.summary = source["summary"];
	        this.methodology = source["methodology"];
	        this.recommendations = source["recommendations"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

