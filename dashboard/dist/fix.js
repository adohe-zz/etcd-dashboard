function showDropdownMenu(ele) {
    var para = ele.parentNode,
        clsName = para.className;
    
    if(clsName === 'co-m-cog ng-isolate-scope') {
        ele.parentNode.className = 'co-m-cog ng-isolate-scope open';
    } else {
        ele.parentNode.className = 'co-m-cog ng-isolate-scope';
    }
}
