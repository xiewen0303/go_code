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
    @Column("{{$v.PropName}}")
    private {{$v.TypeClazz}} {{$v.PropName}};
 {{end}}

}