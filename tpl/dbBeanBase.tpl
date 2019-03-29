package {{.PackageName}};

import java.io.Serializable;
import com.kernel.data.dao.AbsVersion;
import com.kernel.data.dao.IEntity;
import com.kernel.check.db.annotation.*;
import java.sql.*;

@Table("{{.TableName}}")
public class {{.ClazzName}} extends AbsVersion implements Serializable,IEntity{

    @EntityField
	private static final long serialVersionUID = 1L;
 {{range $i,$v :=  .BeanProps}}
    @Column("{{$v.ColumnName}}")
    private {{$v.TypeClazz}} {{$v.PropName}};
 {{end}}

 {{range $i,$v :=  .BeanProps}}
    public {{$v.TypeClazz}} get{{$v.ClazzName}}() {
        return this.{{$v.PropName}};
    }

    public void set{{$v.ClazzName}}({{$v.TypeClazz}} {{$v.PropName}}){
        this.{{$v.PropName}} = {{$v.PropName}};
    }
 {{end}}

    @Override
 	public String getPirmaryKeyName() {
 	{{range $i,$v :=  .BeanProps}}{{if $v.FlagPRI}}
 	    return "{{$v.PropName}}";{{end}}{{end}}
    }

 	@Override
 	public Long getPrimaryKeyValue() {
 	{{range $i,$v :=  .BeanProps}}{{if $v.FlagPRI}}
        return {{$v.PropName}};{{end}}{{end}}
 	}


 	public {{.ClazzName}} copy(){
        {{.ClazzName}} result = new {{.ClazzName}}();{{range $i,$v :=  .BeanProps}}
        result.set{{$v.ClazzName}}(get{{$v.ClazzName}}());{{end}}
    	return result;
    }
}