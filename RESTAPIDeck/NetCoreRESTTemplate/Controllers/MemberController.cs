using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using NetCoreRESTTemplate.Repository;
using NetCoreRESTTemplate.Models;

namespace NetCoreRESTTemplate.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class MemberController : ControllerBase
    {
        private readonly IRepository _repository;
        
        public MemberController(IRepository repository)
        {
            _repository = repository;   
        }
        
        [HttpGet]
        public ActionResult<IEnumerable<Member>> GetAllMembers()
        {
            return Ok(_repository.GetAllMembers());
        }

        [HttpGet("{id}")]
        public ActionResult<Member> GetMember(Guid id)
        {
            return Ok(_repository.GetMember(id));
        }

        [HttpPost]
        public ActionResult<Member> Post([FromBody] MemberRequest memberRequest)
        {
            var member = _repository.CreateMember(memberRequest);
            _repository.AddMember(member);
            return Created(Request.Path.Value, member);
        }

        
        [HttpPut("{id}")]
        public ActionResult<Member> Put(Guid id, [FromBody] MemberRequest memberRequest)
        {
            return Ok(_repository.PutMember(id, memberRequest));
        }

        /// <summary>
        /// Deletes a specific TodoItem.
        /// </summary>
        /// <param name="id"></param>  
        [HttpDelete("{id}")]
        public ActionResult Delete(Guid id)
        {
            _repository.DeleteMember(id);
            return NoContent();
        }
    }
}
