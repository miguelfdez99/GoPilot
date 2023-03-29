export namespace cpu {
	
	export class InfoStat {
	    cpu: number;
	    vendorId: string;
	    family: string;
	    model: string;
	    stepping: number;
	    physicalId: string;
	    coreId: string;
	    cores: number;
	    modelName: string;
	    mhz: number;
	    cacheSize: number;
	    flags: string[];
	    microcode: string;
	
	    static createFrom(source: any = {}) {
	        return new InfoStat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.cpu = source["cpu"];
	        this.vendorId = source["vendorId"];
	        this.family = source["family"];
	        this.model = source["model"];
	        this.stepping = source["stepping"];
	        this.physicalId = source["physicalId"];
	        this.coreId = source["coreId"];
	        this.cores = source["cores"];
	        this.modelName = source["modelName"];
	        this.mhz = source["mhz"];
	        this.cacheSize = source["cacheSize"];
	        this.flags = source["flags"];
	        this.microcode = source["microcode"];
	    }
	}

}

