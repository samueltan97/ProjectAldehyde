using System;
using System.Text;
using System.IO;
using System.Linq;
using System.Collections.Generic;


namespace capitalize
{
    class Program
    {
        static void Main(string[] args)
        {
            string preCaps = File.ReadAllText(args[0], Encoding.UTF8);
            var preCapsArray = StringSeparator(preCaps);
            var postCapsArray = preCapsArray.Select(x => CapitalizeWord(x));
            var postCaps = StringCombiner(postCapsArray);
            File.WriteAllText(args[1], postCaps);
        }

        static List<string> StringSeparator(string originalString) 
        {
            return new List<string>(originalString.Split(new char[0]));
        }

        static string StringCombiner(IEnumerable<string> originalArray) 
        {
            return String.Join(' ', originalArray);
        }

        static string CapitalizeWord(string originalWord)
        {
            return originalWord.Substring(0, 1).ToUpper() + originalWord.Substring(1, originalWord.Length - 1);
        }


    }
}
