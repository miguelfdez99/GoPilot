export namespace backend {
	
	
	
	
	export class MemoryUsage {
	    ram: number;
	    swap: number;
	
	    static createFrom(source: any = {}) {
	        return new MemoryUsage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ram = source["ram"];
	        this.swap = source["swap"];
	    }
	}
	export class NetworkUsage {
	    bytes_sent: number;
	    bytes_received: number;
	
	    static createFrom(source: any = {}) {
	        return new NetworkUsage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.bytes_sent = source["bytes_sent"];
	        this.bytes_received = source["bytes_received"];
	    }
	}
	
	
	
	export class SystemStatThresholds {
	    CPU: number;
	    Memory: number;
	    Disk: number;
	
	    static createFrom(source: any = {}) {
	        return new SystemStatThresholds(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.CPU = source["CPU"];
	        this.Memory = source["Memory"];
	        this.Disk = source["Disk"];
	    }
	}

}

