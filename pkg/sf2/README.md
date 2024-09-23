Run generator
```ps1
go get github.com/zzl/go-tlbimp
go install github.com/zzl/go-tlbimp
# or locally in go-tlbimp repo:
go install .\cmd\...  

go-tlbimp -out-dir .\com\fabrictypes -tlb ".\tlb\FabricTypes.tlb"
go-tlbimp -out-dir .\com\fabriccommon -tlb ".\tlb\FabricCommon.tlb" -imp-pkgs "fabrictypes" -imp-tlbs ".\tlb\FabricTypes.tlb"
go-tlbimp -out-dir .\com\fabricclient -tlb ".\tlb\FabricClient.tlb" -imp-pkgs "fabrictypes;fabriccommon" -imp-tlbs ".\tlb\FabricTypes.tlb;.\tlb\FabricCommon.tlb"
go-tlbimp -out-dir .\com\fabricruntime -tlb ".\tlb\FabricRuntime.tlb" -imp-pkgs "fabrictypes;fabriccommon" -imp-tlbs ".\tlb\FabricTypes.tlb;.\tlb\FabricCommon.tlb"


-imp-pkgs "fabrictypes"
-imp-tlbs ".\tlb\FabricTypes.tlb;.\tlb\FabricCommon.tlb"
```

Ref pack is not well used

go-tlbimp -h
```
 -imp-pkgs string
        import package names(; separated)
  -imp-tlbs string
        import tlb file paths(; separated)
  -out-dir string
        output directory
  -tlb string
        target tlb file path
```

Notes: 
* generator is alpha and generated code frees strings incorrectly.
* CGO on windows not yet supports msvc.
* Golang Syscall fails with ERROR_MOD_NOT_FOUND on free functions.
* In general, golang support for com is not yet mature.