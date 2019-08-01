using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using NetCoreRESTTemplate.Models;

namespace NetCoreRESTTemplate.Repository
{ 
    public class Repository : IRepository 
    {
        private Dictionary<Guid, Member> _repository = new Dictionary<Guid, Member>{ };
        public IEnumerable<Member> GetAllMembers()
        {
            return _repository.Values.ToList();
        }
        public Member GetMember(Guid id)
        {
            return _repository[id];
        }

        public Member CreateMember(MemberRequest memberRequest)
        {
            return new Member(Guid.NewGuid(), memberRequest.Name, memberRequest.Job, memberRequest.Age);
        }

        public Member AddMember(Member member)
        {
            _repository[member.Id] = member;
            return member;
        }

        public Member PutMember(Guid id, MemberRequest memberRequest)
        {
            var member = new Member(id, memberRequest.Name, memberRequest.Job, memberRequest.Age);
            _repository[id] = member;
            return member;
        }

        public void DeleteMember(Guid id)
        {
            _repository.Remove(id);
        }
    }
}
