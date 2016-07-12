
function attachAddRowsEvents() {
    $(".btn.icon-btn.btn-success.education").click(function() {
        $("#educationRows").append(getEducationHtml(getLastIndex("education")));
    });
    $(".btn.icon-btn.btn-success.employment").click(function() {
        $("#employmentRows").append(getEmploymentHtml(getLastIndex("employment")));
    });
    $(".btn.icon-btn.btn-success.skills").click(function() {
        $("#skillsRows").append(getSkillsHtml(getLastIndex("skills")));
    });
    $(".btn.icon-btn.btn-success.facts").click(function() {
        $("#factsRows").append(getFactsHtml(getLastIndex("facts")));
    });
}

function getLastIndex(type) {
    var lastNameId = $(".form-control."+type+"-first").last().attr('id');
    var index = 0;

    if (typeof  lastNameId != 'undefined') {
        index = parseInt(lastNameId.replace(/[^0-9]/gi, ''));
        if ($.isNumeric(index))
            index ++;
    }

    return index;
}

function getEducationHtml(index) {
    return '<div class="form-inline text-center"> \
            <div class="form-group"> \
                <label for="inputEducationFrom'+index+'"></label> \
                <div class="col-sm-2"> \
                    <input type="number" class="form-control education-first" id="inputEducationFrom'+index+'" name="Education.From" placeholder="2001" value=""> \
                </div> \
            </div> \
            <div class="form-group"> \
                <label for="inputEducationTo'+index+'"></label> \
                <div class="col-sm-2"> \
                    <input type="text" class="form-control" id="inputEducationTo'+index+'" name="Education.To" placeholder="2016" value=""> \
                </div> \
            </div> \
            <div class="form-group"> \
                <label for="inputEducationGrade'+index+'"></label> \
                <div class="col-sm-2"> \
                    <input type="text" class="form-control" id="inputEducationGrade'+index+'" name="Education.Grade" placeholder="Master Degree" value=""> \
                </div> \
            </div> \
            <div class="form-group"> \
                <label for="inputEducationScience'+index+'"></label> \
                <div class="col-sm-2"> \
                    <input type="text" class="form-control" id="inputEducationScience'+index+'" name="Education.Science" placeholder="Computer Science" value=""> \
                </div> \
            </div> \
            <div class="form-group"> \
                <label for="inputEducationInstitution'+index+'"></label> \
                <div class="col-sm-2"> \
                    <input type="text" class="form-control" id="inputEducationInstitution'+index+'" name="Education.Institution" placeholder="The Harvard Academy" value=""> \
                </div> \
            </div> \
            <div class="form-group"> \
                <label for="inputEducationDescription'+index+'"></label> \
                <div class="col-sm-2"> \
                    <input type="text" class="form-control" id="inputEducationDescription'+index+'" name="Education.Description" placeholder="The best education" value=""> \
                </div> \
            </div> \
            ' +getRemoveBtnHtml()+ ' \
        </div> \
    ';
}

function getEmploymentHtml(index) {
    return '<div class="form-inline text-center"> \
            <div class="form-group"> \
                <label for="inputEmploymentFrom'+index+'"></label> \
                <div class="col-sm-2"> \
                    <input type="number" class="form-control education-first" id="inputEmploymentFrom'+index+'" name="Employment.From" placeholder="2001" value=""> \
                </div> \
            </div> \
            <div class="form-group"> \
                <label for="inputEmploymentTo'+index+'"></label> \
                <div class="col-sm-2"> \
                    <input type="text" class="form-control" id="inputEmploymentTo'+index+'" name="Employment.To" placeholder="2016" value=""> \
                </div> \
            </div> \
            <div class="form-group"> \
                <label for="inputEmploymentCompany'+index+'"></label> \
                <div class="col-sm-2"> \
                    <input type="text" class="form-control" id="inputEmploymentCompany'+index+'" name="Employment.Company" placeholder="Google" value=""> \
                </div> \
            </div> \
            <div class="form-group"> \
                <label for="inputEmploymentPosition'+index+'"></label> \
                <div class="col-sm-2"> \
                    <input type="text" class="form-control" id="inputEmploymentPosition'+index+'" name="Employment.Position" placeholder="Boss" value=""> \
                </div> \
            </div> \
            <div class="form-group"> \
                <label for="inputEmploymentDescription'+index+'"></label> \
                <div class="col-sm-2"> \
                    <input type="text" class="form-control" id="inputEmploymentDescription'+index+'" name="Employment.Description" placeholder="The best developer" value=""> \
                </div> \
            </div> \
            ' +getRemoveBtnHtml()+ ' \
        </div> \
    ';
}

function getSkillsHtml(index) {
    return '<div class="form-inline text-center"> \
            <div class="form-group"> \
                <label for="inputSkillName'+index+'"></label> \
                <div class="col-sm-2"> \
                    <input type="text" class="form-control skills-first" id="inputSkillName'+index+'" name="Skill.Name" placeholder="PHP" value=""> \
                </div> \
            </div> \
            <div class="form-group"> \
                <label for="inputSkillLevelInPct'+index+'"></label> \
                <div class="col-sm-2"> \
                    <input type="number" class="form-control" id="inputSkillLevelInPct'+index+'" name="Skill.LevelInPct" placeholder="100" value=""> \
                </div> \
            </div> \
            <div class="form-group"> \
                <label for="inputSkillColor'+index+'"></label> \
                <div class="col-sm-2"> \
                    <input type="text" class="form-control" id="inputSkillColor'+index+'" name="Skill.Color" placeholder="red" value=""> \
                </div> \
            </div> \
            ' +getRemoveBtnHtml()+ ' \
        </div> \
    ';
}

function getFactsHtml(index) {
    return '<div class="form-inline text-center"> \
            <div class="form-group"> \
                <label for="inputFactName'+index+'"></label> \
                <div class="col-sm-2"> \
                    <input type="text" class="form-control facts-first" id="inputFactName'+index+'" name="Fact.Name" placeholder="Cup of tea" value=""> \
                </div> \
            </div> \
            <div class="form-group"> \
                <label for="inputFactValue'+index+'"></label> \
                <div class="col-sm-2"> \
                    <input type="text" class="form-control" id="inputFactValue'+index+'" name="Fact.Value" placeholder="1000" value=""> \
                </div> \
            </div> \
            ' +getRemoveBtnHtml()+ ' \
        </div> \
    ';
}

function getRemoveBtnHtml() {
    return '<div class="form-group"> \
                <div class="col-sm-offset-2 col-sm-5"> \
                    <a class="btn icon-btn btn-danger" href="#" onclick="removeRow(this); return false;"> \
                        <span class="glyphicon btn-glyphicon glyphicon-trash img-circle text-danger"></span> \
                        Delete \
                    </a> \
                </div> \
            </div> \
        ';
}

function removeRow(_button) {
    $(_button).parent().parent().parent().remove();
}
