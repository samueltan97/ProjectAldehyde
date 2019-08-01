using System;
using System.Collections.Generic;
using NetCoreRESTTemplate.Models;

namespace NetCoreRESTTemplate.Repository
{
    public interface IRepository
    { 
        IEnumerable<Member> GetAllMembers();
        Member GetMember(Guid id);
        Member CreateMember(MemberRequest memberRequest);
        Member AddMember(Member member);
        Member PutMember(Guid id, MemberRequest memberRequest);
        void DeleteMember(Guid id);
    }

}
